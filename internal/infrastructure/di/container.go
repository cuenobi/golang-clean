package di

import (
	"github.com/IBM/sarama"
	orderusecase "github.com/cuenobi/golang-clean/internal/application/usecase/order"
	userusecase "github.com/cuenobi/golang-clean/internal/application/usecase/user"
	messaginginfra "github.com/cuenobi/golang-clean/internal/infrastructure/messaging"
	persistenceinfra "github.com/cuenobi/golang-clean/internal/infrastructure/persistence"
	messageadapter "github.com/cuenobi/golang-clean/internal/interfaces/messaging"
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/cuenobi/golang-clean/internal/shared/logger"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Container is the composition root for dependency wiring.
// It centralizes runtime dependencies and shared service lifecycles.
type Container struct {
	Cfg config.Config
	DB  *gorm.DB

	Log   logger.Logger
	Clock outClock
	IDGen outIDGenerator
	Tx    kernel.TxManager

	OrderRepo  *persistenceinfra.OrderRepository
	UserRepo   *persistenceinfra.UserRepository
	OutboxRepo *persistenceinfra.OutboxRepository

	Producer       sarama.SyncProducer
	EventPublisher *messaginginfra.Publisher

	OrderUseCase *orderusecase.OrderUseCase
	UserUseCase  *userusecase.UserUseCase

	HTTPApp  *fiber.App
	Consumer *messageadapter.Consumer
}

// local aliases make intent explicit while keeping container fields concise.
type (
	outClock       = kernel.SystemClock
	outIDGenerator = kernel.UUIDGenerator
)

func NewContainer(cfg config.Config, db *gorm.DB) (*Container, error) {
	c := &Container{
		Cfg: cfg,
		DB:  db,
	}

	c.wireCore()
	c.wirePersistence()

	if err := c.wireMessaging(); err != nil {
		return nil, err
	}

	c.wireUseCases()
	c.wireHTTP()
	c.wireConsumer()

	return c, nil
}

func (c *Container) wireCore() {
	c.Log = logger.New(c.Cfg)
	c.Clock = kernel.SystemClock{}
	c.IDGen = kernel.UUIDGenerator{}
	c.Tx = sharedpersistence.NewGormTxManager(c.DB)
}

func (c *Container) wirePersistence() {
	c.OrderRepo = persistenceinfra.NewOrderRepository(c.DB)
	c.UserRepo = persistenceinfra.NewUserRepository(c.DB)
	c.OutboxRepo = persistenceinfra.NewOutboxRepository(c.DB)
}

func (c *Container) wireUseCases() {
	c.OrderUseCase = orderusecase.NewOrderUseCase(c.OrderRepo, c.Tx, c.OutboxRepo, c.Clock, c.IDGen)
	c.UserUseCase = userusecase.NewUserUseCase(c.UserRepo, c.Clock, c.IDGen)
}

func (c *Container) wireConsumer() {
	c.Consumer = messageadapter.NewConsumer(c.Log, c.OutboxRepo, c.EventPublisher, c.Cfg, c.Clock)
}

func (c *Container) Close() {
	if c.Producer != nil {
		_ = c.Producer.Close()
	}
}
