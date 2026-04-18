package http

type OrderResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
}
