package usecase

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
)

func (u *OrderUseCase) GetOrder(ctx context.Context, id string) (dto.OrderResponse, error) {
	order, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	return toOrderResponse(order), nil
}
