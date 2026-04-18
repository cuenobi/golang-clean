package http

import "github.com/cuenobi/golang-clean/internal/user/application/dto"

func toCreateUserDTO(req CreateUserRequest) dto.CreateUserRequest {
	return dto.CreateUserRequest{Name: req.Name, Email: req.Email}
}

func toUpdateUserDTO(req UpdateUserRequest) dto.UpdateUserRequest {
	return dto.UpdateUserRequest{Name: req.Name, Email: req.Email}
}

func toUserResponse(resp dto.UserResponse) UserResponse {
	return UserResponse{ID: resp.ID, Name: resp.Name, Email: resp.Email}
}
