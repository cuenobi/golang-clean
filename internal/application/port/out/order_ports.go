package out

import (
	"context"
	"time"

	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/event"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

type OrderRepository interface {
	Save(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, orderID string) (*entity.Order, error)
	GetByIdempotencyKey(ctx context.Context, idempotencyKey string) (*entity.Order, error)
	List(ctx context.Context) ([]*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, orderID string) error
}

type TxManager = kernel.TxManager

type EventPublisher interface {
	PublishOrderCreated(ctx context.Context, eventPayload any) error
}

type OrderEventOutboxWriter interface {
	EnqueueOrderCreated(ctx context.Context, event event.OrderCreated) error
}

type OutboxMessage struct {
	ID         string
	EventType  string
	Payload    []byte
	RetryCount int
}

type OutboxStore interface {
	ClaimPending(ctx context.Context, now time.Time, limit int, processingTimeout time.Duration) ([]OutboxMessage, error)
	MarkPublished(ctx context.Context, messageID string, publishedAt time.Time) error
	MarkRetry(ctx context.Context, messageID string, retryCount int, nextRetryAt time.Time, lastError string, dead bool) error
}

type Clock = kernel.Clock
type IDGenerator = kernel.IDGenerator
