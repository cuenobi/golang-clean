package di

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/cuenobi/golang-clean/internal/order/application/usecase"
	messaginginfra "github.com/cuenobi/golang-clean/internal/order/infrastructure/messaging"
	persistenceinfra "github.com/cuenobi/golang-clean/internal/order/infrastructure/persistence"
	httpadapter "github.com/cuenobi/golang-clean/internal/order/interfaces/http"
	messageadapter "github.com/cuenobi/golang-clean/internal/order/interfaces/messaging"
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/cuenobi/golang-clean/internal/shared/kafkax"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/cuenobi/golang-clean/internal/shared/logger"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	userusecase "github.com/cuenobi/golang-clean/internal/user/application/usecase"
	userpersistence "github.com/cuenobi/golang-clean/internal/user/infrastructure/persistence"
	userhttp "github.com/cuenobi/golang-clean/internal/user/interfaces/http"
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
	tx := sharedpersistence.NewGormTxManager(db)
	kafkaConfig := kafkax.NewDefaultSaramaConfig(cfg.AppName)
	producer, err := sarama.NewSyncProducer(cfg.KafkaBrokers, kafkaConfig)
	if err != nil {
		return nil, err
	}
	publisher := messaginginfra.NewPublisher(producer, "order.created.v1")
	idgen := kernel.UUIDGenerator{}
	useCase := usecase.NewOrderUseCase(repo, tx, publisher, kernel.SystemClock{}, idgen)
	userRepo := userpersistence.NewUserRepository(db)
	userUC := userusecase.NewUserUseCase(userRepo, kernel.SystemClock{}, idgen)

	app := fiber.New(fiber.Config{ErrorHandler: httpx.NewErrorHandler()})
	app.Use(httpx.RequestIDMiddleware())
	app.Use(httpx.RequestLogger(log))
	v1 := app.Group("/api/v1")
	httpHandler := httpadapter.NewHandler(useCase)
	httpadapter.RegisterRoutes(v1, httpHandler)
	userHandler := userhttp.NewHandler(userUC)
	userhttp.RegisterRoutes(v1, userHandler)

	consumer := messageadapter.NewConsumer(log)

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
