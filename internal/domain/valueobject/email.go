package valueobject

import (
	"fmt"
	"net/mail"
	"strings"
)

type Email string

func NewEmail(raw string) (Email, error) {
	trimmed := strings.TrimSpace(strings.ToLower(raw))
	if trimmed == "" {
		return "", fmt.Errorf("email is required")
	}
	if _, err := mail.ParseAddress(trimmed); err != nil {
		return "", fmt.Errorf("invalid email format")
	}
	return Email(trimmed), nil
}
