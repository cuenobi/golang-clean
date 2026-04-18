package di

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/cuenobi/golang-clean/internal/application/usecase"
	messaginginfra "github.com/cuenobi/golang-clean/internal/infrastructure/messaging"
	persistenceinfra "github.com/cuenobi/golang-clean/internal/infrastructure/persistence"
	httpadapter "github.com/cuenobi/golang-clean/internal/interfaces/http/order"
	systemhttp "github.com/cuenobi/golang-clean/internal/interfaces/http/system"
	userhttp "github.com/cuenobi/golang-clean/internal/interfaces/http/user"
	messageadapter "github.com/cuenobi/golang-clean/internal/interfaces/messaging"
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/cuenobi/golang-clean/internal/shared/kafkax"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/cuenobi/golang-clean/internal/shared/logger"
	"github.com/cuenobi/golang-clean/internal/shared/metrics"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Module struct {
	HTTPRunner     func() error
	ConsumerRunner func(context.Context) error
}

func NewModule(cfg config.Config, db *gorm.DB) (*Module, error) {
	log := logger.New(cfg)
	repo := persistenceinfra.NewOrderRepository(db)
	outboxRepo := persistenceinfra.NewOutboxRepository(db)
	tx := sharedpersistence.NewGormTxManager(db)
	kafkaConfig := kafkax.NewDefaultSaramaConfig(cfg)
	producer, err := sarama.NewSyncProducer(cfg.KafkaBrokers, kafkaConfig)
	if err != nil {
		return nil, err
	}
	publisher := messaginginfra.NewPublisher(producer, "order.created.v1")
	clock := kernel.SystemClock{}
	idgen := kernel.UUIDGenerator{}
	useCase := usecase.NewOrderUseCase(repo, tx, outboxRepo, clock, idgen)
	userRepo := persistenceinfra.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo, clock, idgen)

	readTimeout := time.Duration(cfg.HTTPReadTimeout) * time.Second
	writeTimeout := time.Duration(cfg.HTTPWriteTimeout) * time.Second
	app := fiber.New(fiber.Config{
		ErrorHandler:  httpx.NewErrorHandler(),
		ReadTimeout:   readTimeout,
		WriteTimeout:  writeTimeout,
		CaseSensitive: true,
	})
	app.Use(httpx.RequestIDMiddleware())
	app.Use(metrics.HTTPMiddleware())
	app.Use(httpx.RateLimiter(cfg))
	app.Use(httpx.RequestLogger(log))
	systemHandler := systemhttp.NewHandler(cfg, db)
	systemhttp.RegisterRoutes(app, systemHandler)

	v1 := app.Group("/api/v1")
	httpHandler := httpadapter.NewHandler(useCase)
	httpadapter.RegisterRoutes(v1, httpHandler, cfg)
	userHandler := userhttp.NewHandler(userUC)
	userhttp.RegisterRoutes(v1, userHandler, cfg)

	consumer := messageadapter.NewConsumer(log, outboxRepo, publisher, cfg, clock)

	return &Module{
		HTTPRunner: func() error {
			defer producer.Close()
			return app.Listen(cfg.HTTPAddress)
		},
		ConsumerRunner: func(ctx context.Context) error {
			defer producer.Close()
			return consumer.Run(ctx)
		},
	}, nil
}
