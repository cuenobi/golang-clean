package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/cuenobi/golang-clean/internal/user/domain/valueobject"
)

type User struct {
	ID        string
	Name      string
	Email     valueobject.Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id, name string, email valueobject.Email, now time.Time) (*User, error) {
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		return nil, fmt.Errorf("name is required")
	}
	return &User{
		ID:        id,
		Name:      cleanName,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) Update(name string, email valueobject.Email, now time.Time) error {
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		return fmt.Errorf("name is required")
	}
	u.Name = cleanName
	u.Email = email
	u.UpdatedAt = now
	return nil
}
