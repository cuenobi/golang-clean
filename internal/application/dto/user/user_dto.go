package userdto

type CreateUserRequest struct {
	Name  string
	Email string
}

type UpdateUserRequest struct {
	Name  string
	Email string
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
