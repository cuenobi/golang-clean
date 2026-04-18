package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/event"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

func (u *OrderUseCase) CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (dto.OrderResponse, error) {
	idempotencyKey := strings.TrimSpace(req.IdempotencyKey)
	if idempotencyKey != "" {
		existing, err := u.repo.GetByIdempotencyKey(ctx, idempotencyKey)
		switch {
		case err == nil && existing != nil:
			return toOrderResponse(existing), nil
		case err == nil && existing == nil:
			// continue
		case errors.Is(err, kernel.ErrNotFound):
			// continue
		default:
			return dto.OrderResponse{}, err
		}
	}

	money, err := valueobject.NewMoney(req.Currency, req.Amount)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	var created *entity.Order
	if err := u.tx.WithinTransaction(ctx, func(txCtx context.Context) error {
		order, err := entity.NewOrder(u.idGen.NewID(), req.CustomerID, idempotencyKey, money, u.clock.Now())
		if err != nil {
			return err
		}
		if err := u.repo.Save(txCtx, order); err != nil {
			return err
		}

		outboxEvent := event.OrderCreated{
			OrderID:    order.ID,
			CustomerID: order.CustomerID,
			Currency:   order.Amount.Currency,
			Amount:     order.Amount.Amount,
		}
		if err := u.outbox.EnqueueOrderCreated(txCtx, outboxEvent); err != nil {
			return fmt.Errorf("enqueue order created outbox event: %w", err)
		}
		created = order
		return nil
	}); err != nil {
		return dto.OrderResponse{}, err
	}

	if created == nil {
		return dto.OrderResponse{}, fmt.Errorf("create order failed")
	}
	return toOrderResponse(created), nil
}
