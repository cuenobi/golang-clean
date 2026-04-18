package persistence

import "time"

type OrderModel struct {
	ID             string    `gorm:"column:id;primaryKey"`
	CustomerID     string    `gorm:"column:customer_id;not null"`
	IdempotencyKey *string   `gorm:"column:idempotency_key;uniqueIndex"`
	Currency       string    `gorm:"column:currency;not null"`
	Amount         int64     `gorm:"column:amount;not null"`
	Status         string    `gorm:"column:status;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null"`
}

func (OrderModel) TableName() string {
	return "orders"
}
