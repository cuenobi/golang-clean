package usecase

import (
	"context"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
)

func (u *UserUseCase) GetUser(ctx context.Context, id string) (dto.UserResponse, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return toUserResponse(user), nil
}
