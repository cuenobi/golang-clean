package usecase

import (
	"context"
	"errors"
	"fmt"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

func (u *UserUseCase) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	existing, err := u.repo.GetByEmail(ctx, string(email))
	if err != nil && !errors.Is(err, kernel.ErrNotFound) {
		return dto.UserResponse{}, err
	}
	if existing != nil {
		return dto.UserResponse{}, fmt.Errorf("%w: email already exists", kernel.ErrConflict)
	}

	created, err := entity.NewUser(u.idGen.NewID(), req.Name, email, u.clock.Now())
	if err != nil {
		return dto.UserResponse{}, err
	}
	if err := u.repo.Create(ctx, created); err != nil {
		return dto.UserResponse{}, err
	}
	return toUserResponse(created), nil
}
