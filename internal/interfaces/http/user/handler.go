package http

import (
	"errors"

	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedvalidator "github.com/cuenobi/golang-clean/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase in.UserUseCase
}

func NewHandler(useCase in.UserUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.useCase.CreateUser(c.UserContext(), toCreateUserDTO(req))
	if err != nil {
		if errors.Is(err, kernel.ErrConflict) {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(toUserResponse(result))
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	result, err := h.useCase.GetUser(c.UserContext(), c.Params("id"))
	if err != nil {
		if errors.Is(err, kernel.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}
	return c.JSON(toUserResponse(result))
}

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

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.useCase.UpdateUser(c.UserContext(), c.Params("id"), toUpdateUserDTO(req))
	if err != nil {
		if errors.Is(err, kernel.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		if errors.Is(err, kernel.ErrConflict) {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return err
	}
	return c.JSON(toUserResponse(result))
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	err := h.useCase.DeleteUser(c.UserContext(), c.Params("id"))
	if err != nil {
		if errors.Is(err, kernel.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
