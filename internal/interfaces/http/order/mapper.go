package http

import dto "github.com/cuenobi/golang-clean/internal/application/dto/order"

func toCreateOrderDTO(req CreateOrderRequest, idempotencyKey string) dto.CreateOrderRequest {
	return dto.CreateOrderRequest{
		CustomerID:     req.CustomerID,
		Currency:       req.Currency,
		Amount:         req.Amount,
		IdempotencyKey: idempotencyKey,
	}
}

func toUpdateOrderDTO(req UpdateOrderRequest) dto.UpdateOrderRequest {
	return dto.UpdateOrderRequest{
		CustomerID: req.CustomerID,
		Currency:   req.Currency,
		Amount:     req.Amount,
	}
}

func toOrderResponse(resp dto.OrderResponse) OrderResponse {
	return OrderResponse{
		ID:         resp.ID,
		CustomerID: resp.CustomerID,
		Currency:   resp.Currency,
		Amount:     resp.Amount,
		Status:     resp.Status,
	}
}
