package dto

type CreateOrderRequest struct {
	CustomerID     string
	Currency       string
	Amount         int64
	IdempotencyKey string
}

type UpdateOrderRequest struct {
	CustomerID string
	Currency   string
	Amount     int64
}

type OrderResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
}
