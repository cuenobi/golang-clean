package usecase

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
)

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
