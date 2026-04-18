package valueobject

import "fmt"

type Money struct {
	Currency string
	Amount   int64
}

func NewMoney(currency string, amount int64) (Money, error) {
	if currency == "" {
		return Money{}, fmt.Errorf("currency is required")
	}
	if amount <= 0 {
		return Money{}, fmt.Errorf("amount must be positive")
	}
	return Money{Currency: currency, Amount: amount}, nil
}
