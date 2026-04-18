package event

type OrderCreated struct {
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"`
}
