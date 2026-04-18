package usecase

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
)

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
