package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
)

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusConfirmed Status = "CONFIRMED"
	StatusCanceled  Status = "CANCELED"
)

type Order struct {
	ID             string
	CustomerID     string
	IdempotencyKey string
	Amount         valueobject.Money
	Status         Status
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewOrder(id, customerID, idempotencyKey string, amount valueobject.Money, now time.Time) (*Order, error) {
	if strings.TrimSpace(customerID) == "" {
		return nil, fmt.Errorf("customer id is required")
	}
	return &Order{
		ID:             id,
		CustomerID:     strings.TrimSpace(customerID),
		IdempotencyKey: strings.TrimSpace(idempotencyKey),
		Amount:         amount,
		Status:         StatusPending,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (o *Order) Update(customerID string, amount valueobject.Money, now time.Time) error {
	if strings.TrimSpace(customerID) == "" {
		return fmt.Errorf("customer id is required")
	}
	o.CustomerID = strings.TrimSpace(customerID)
	o.Amount = amount
	o.UpdatedAt = now
	return nil
}

func (o *Order) Confirm(now time.Time) error {
	if o.Status != StatusPending {
		return fmt.Errorf("order can only be confirmed from pending status")
	}
	o.Status = StatusConfirmed
	o.UpdatedAt = now
	return nil
}

func (o *Order) Cancel(now time.Time) error {
	if o.Status == StatusCanceled {
		return fmt.Errorf("order already canceled")
	}
	o.Status = StatusCanceled
	o.UpdatedAt = now
	return nil
}
