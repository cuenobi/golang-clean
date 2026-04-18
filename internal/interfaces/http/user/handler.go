package http

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	sharedhttpx "github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedvalidator "github.com/cuenobi/golang-clean/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase in.UserUseCase
}

type ErrorResponseDoc = sharedhttpx.ErrorResponse

func NewHandler(useCase in.UserUseCase) *Handler {
	return &Handler{useCase: useCase}
}

// CreateUser godoc
// @Summary Create user
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param X-Permissions header string true "Required permission: users:write"
// @Param request body CreateUserRequest true "Create user payload"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 409 {object} ErrorResponseDoc
// @Failure 422 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return kernel.NewBadRequestError("invalid request body")
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return err
	}

	result, err := h.useCase.CreateUser(c.UserContext(), toCreateUserDTO(req))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(toUserResponse(result))
}

// GetUser godoc
// @Summary Get user by ID
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: users:read"
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	result, err := h.useCase.GetUser(c.UserContext(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(toUserResponse(result))
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: users:read"
// @Success 200 {array} UserResponse
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/users [get]
func (h *Handler) ListUsers(c *fiber.Ctx) error {
	result, err := h.useCase.ListUsers(c.UserContext())
	if err != nil {
		return err
	}
	items := make([]UserResponse, 0, len(result))
	for _, user := range result {
		items = append(items, toUserResponse(user))
	}
	return c.JSON(items)
}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param X-Permissions header string true "Required permission: users:write"
// @Param id path string true "User ID"
// @Param request body UpdateUserRequest true "Update user payload"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 409 {object} ErrorResponseDoc
// @Failure 422 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/users/{id} [put]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return kernel.NewBadRequestError("invalid request body")
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return err
	}

	result, err := h.useCase.UpdateUser(c.UserContext(), c.Params("id"), toUpdateUserDTO(req))
	if err != nil {
		return err
	}
	return c.JSON(toUserResponse(result))
}

// DeleteUser godoc
// @Summary Delete user
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: users:write"
// @Param id path string true "User ID"
// @Success 204
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	if err := h.useCase.DeleteUser(c.UserContext(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
