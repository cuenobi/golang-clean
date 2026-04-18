package usecase

import (
	"context"
	"errors"
	"fmt"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

func (u *UserUseCase) UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	existing, err := u.repo.GetByEmail(ctx, string(email))
	if err != nil && !errors.Is(err, kernel.ErrNotFound) {
		return dto.UserResponse{}, err
	}
	if existing != nil && existing.ID != id {
		return dto.UserResponse{}, fmt.Errorf("%w: email already exists", kernel.ErrConflict)
	}

	if err := user.Update(req.Name, email, u.clock.Now()); err != nil {
		return dto.UserResponse{}, err
	}
	if err := u.repo.Update(ctx, user); err != nil {
		return dto.UserResponse{}, err
	}
	return toUserResponse(user), nil
}
