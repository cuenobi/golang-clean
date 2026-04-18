package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/cuenobi/golang-clean/internal/application/dto"
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

var _ in.UserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	repo  out.UserRepository
	clock out.Clock
	idGen out.IDGenerator
}

func NewUserUseCase(repo out.UserRepository, clock out.Clock, idGen out.IDGenerator) *UserUseCase {
	return &UserUseCase{repo: repo, clock: clock, idGen: idGen}
}

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

func (u *UserUseCase) GetUser(ctx context.Context, id string) (dto.UserResponse, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return toUserResponse(user), nil
}

func (u *UserUseCase) ListUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		result = append(result, toUserResponse(user))
	}
	return result, nil
}

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

func (u *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func toUserResponse(user *entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: string(user.Email),
	}
}
