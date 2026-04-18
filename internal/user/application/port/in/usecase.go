package in

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/user/application/dto"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
	GetUser(ctx context.Context, id string) (dto.UserResponse, error)
	ListUsers(ctx context.Context) ([]dto.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
}
