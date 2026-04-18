package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cuenobi/golang-clean/internal/application/dto"
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/event"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

var _ in.OrderUseCase = (*OrderUseCase)(nil)

type OrderUseCase struct {
	repo   out.OrderRepository
	tx     out.TxManager
	outbox out.OrderEventOutboxWriter
	clock  out.Clock
	idGen  out.IDGenerator
}

func NewOrderUseCase(
	repo out.OrderRepository,
	tx out.TxManager,
	outbox out.OrderEventOutboxWriter,
	clock out.Clock,
	idGen out.IDGenerator,
) *OrderUseCase {
	return &OrderUseCase{
		repo:   repo,
		tx:     tx,
		outbox: outbox,
		clock:  clock,
		idGen:  idGen,
	}
}

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

func (u *OrderUseCase) GetOrder(ctx context.Context, id string) (dto.OrderResponse, error) {
	order, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	return toOrderResponse(order), nil
}

func (u *OrderUseCase) ListOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	orders, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]dto.OrderResponse, 0, len(orders))
	for _, order := range orders {
		result = append(result, toOrderResponse(order))
	}
	return result, nil
}

func (u *OrderUseCase) UpdateOrder(ctx context.Context, id string, req dto.UpdateOrderRequest) (dto.OrderResponse, error) {
	order, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	money, err := valueobject.NewMoney(req.Currency, req.Amount)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	if err := order.Update(req.CustomerID, money, u.clock.Now()); err != nil {
		return dto.OrderResponse{}, err
	}
	if err := u.repo.Update(ctx, order); err != nil {
		return dto.OrderResponse{}, err
	}
	return toOrderResponse(order), nil
}

func (u *OrderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func toOrderResponse(order *entity.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Currency:   order.Amount.Currency,
		Amount:     order.Amount.Amount,
		Status:     string(order.Status),
	}
}
