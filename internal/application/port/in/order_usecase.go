package in

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (dto.OrderResponse, error)
	GetOrder(ctx context.Context, id string) (dto.OrderResponse, error)
	ListOrders(ctx context.Context) ([]dto.OrderResponse, error)
	UpdateOrder(ctx context.Context, id string, req dto.UpdateOrderRequest) (dto.OrderResponse, error)
	DeleteOrder(ctx context.Context, id string) error
}
