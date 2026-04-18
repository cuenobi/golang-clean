package persistence

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/event"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const orderCreatedEventType = "order.created.v1"

var (
	_ out.OrderEventOutboxWriter = (*OutboxRepository)(nil)
	_ out.OutboxStore            = (*OutboxRepository)(nil)
)

type OutboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) *OutboxRepository {
	return &OutboxRepository{db: db}
}

func (r *OutboxRepository) EnqueueOrderCreated(ctx context.Context, evt event.OrderCreated) error {
	payload, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	model := OutboxMessageModel{
		ID:            uuid.NewString(),
		AggregateType: "order",
		AggregateID:   evt.OrderID,
		EventType:     orderCreatedEventType,
		Payload:       payload,
		Status:        OutboxStatusPending,
		RetryCount:    0,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Create(&model).Error
}

func (r *OutboxRepository) ClaimPending(
	ctx context.Context,
	now time.Time,
	limit int,
	processingTimeout time.Duration,
) ([]out.OutboxMessage, error) {
	now = now.UTC()

	if limit <= 0 {
		limit = 50
	}

	staleProcessingAt := now.Add(-processingTimeout)
	if processingTimeout <= 0 {
		staleProcessingAt = now.Add(-15 * time.Second)
	}

	rows := make([]OutboxMessageModel, 0, limit)
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := `
WITH candidates AS (
  SELECT id
  FROM outbox_messages
  WHERE
    (
      status = ? AND (next_retry_at IS NULL OR next_retry_at <= ?)
    )
    OR
    (
      status = ? AND updated_at <= ?
    )
  ORDER BY created_at ASC
  LIMIT ?
  FOR UPDATE SKIP LOCKED
)
UPDATE outbox_messages o
SET status = ?, updated_at = ?
FROM candidates
WHERE o.id = candidates.id
RETURNING o.id, o.event_type, o.payload, o.retry_count;
`
		return tx.Raw(
			query,
			OutboxStatusPending,
			now,
			OutboxStatusProcessing,
			staleProcessingAt,
			limit,
			OutboxStatusProcessing,
			now,
		).Scan(&rows).Error
	})
	if err != nil {
		return nil, err
	}

	result := make([]out.OutboxMessage, 0, len(rows))
	for _, row := range rows {
		result = append(result, out.OutboxMessage{
			ID:         row.ID,
			EventType:  row.EventType,
			Payload:    row.Payload,
			RetryCount: row.RetryCount,
		})
	}
	return result, nil
}

func (r *OutboxRepository) MarkPublished(ctx context.Context, messageID string, publishedAt time.Time) error {
	publishedAt = publishedAt.UTC()

	return sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).
		Model(&OutboxMessageModel{}).
		Where("id = ?", messageID).
		Updates(map[string]any{
			"status":       OutboxStatusPublished,
			"published_at": publishedAt,
			"updated_at":   publishedAt,
			"last_error":   nil,
		}).Error
}

func (r *OutboxRepository) MarkRetry(
	ctx context.Context,
	messageID string,
	retryCount int,
	nextRetryAt time.Time,
	lastError string,
	dead bool,
) error {
	nextRetryAt = nextRetryAt.UTC()

	status := OutboxStatusPending
	if dead {
		status = OutboxStatusDead
	}

	lastErr := lastError
	updates := map[string]any{
		"status":        status,
		"retry_count":   retryCount,
		"next_retry_at": nextRetryAt,
		"last_error":    &lastErr,
		"updated_at":    time.Now().UTC(),
	}

	return sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).
		Model(&OutboxMessageModel{}).
		Where("id = ?", messageID).
		Updates(updates).Error
}
