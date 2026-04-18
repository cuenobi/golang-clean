package out

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

type OrderRepository interface {
	Save(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, orderID string) (*entity.Order, error)
	List(ctx context.Context) ([]*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, orderID string) error
}

type TxManager = kernel.TxManager

type EventPublisher interface {
	PublishOrderCreated(ctx context.Context, eventPayload any) error
}

type Clock = kernel.Clock
type IDGenerator = kernel.IDGenerator
