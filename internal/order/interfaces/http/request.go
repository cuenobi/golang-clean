package http

type CreateOrderRequest struct {
	CustomerID string `json:"customer_id" validate:"required,min=2,max=100"`
	Currency   string `json:"currency" validate:"required,len=3"`
	Amount     int64  `json:"amount" validate:"required,gt=0"`
}

type UpdateOrderRequest struct {
	CustomerID string `json:"customer_id" validate:"required,min=2,max=100"`
	Currency   string `json:"currency" validate:"required,len=3"`
	Amount     int64  `json:"amount" validate:"required,gt=0"`
}
