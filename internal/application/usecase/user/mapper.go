package usecase

import (
	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
)

func toUserResponse(user *entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: string(user.Email),
	}
}
