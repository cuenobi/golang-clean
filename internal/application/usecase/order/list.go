package usecase

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
)

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
