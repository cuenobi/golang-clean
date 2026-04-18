package usecase

import (
	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
)

func toOrderResponse(order *entity.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Currency:   order.Amount.Currency,
		Amount:     order.Amount.Amount,
		Status:     string(order.Status),
	}
}
