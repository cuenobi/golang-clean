package persistence

import "time"

const (
	OutboxStatusPending    = "PENDING"
	OutboxStatusProcessing = "PROCESSING"
	OutboxStatusPublished  = "PUBLISHED"
	OutboxStatusDead       = "DEAD"
)

type OutboxMessageModel struct {
	ID            string     `gorm:"column:id;primaryKey"`
	AggregateType string     `gorm:"column:aggregate_type;not null"`
	AggregateID   string     `gorm:"column:aggregate_id;not null"`
	EventType     string     `gorm:"column:event_type;not null"`
	Payload       []byte     `gorm:"column:payload;not null"`
	Status        string     `gorm:"column:status;not null"`
	RetryCount    int        `gorm:"column:retry_count;not null"`
	NextRetryAt   *time.Time `gorm:"column:next_retry_at"`
	LastError     *string    `gorm:"column:last_error"`
	PublishedAt   *time.Time `gorm:"column:published_at"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;not null"`
}

func (OutboxMessageModel) TableName() string {
	return "outbox_messages"
}
