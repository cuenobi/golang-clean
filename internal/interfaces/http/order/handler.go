package http

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedvalidator "github.com/cuenobi/golang-clean/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase in.OrderUseCase
}

func NewHandler(useCase in.OrderUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var req CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return kernel.NewBadRequestError("invalid request body")
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return err
	}

	result, err := h.useCase.CreateOrder(c.UserContext(), toCreateOrderDTO(req, c.Get("Idempotency-Key")))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(toOrderResponse(result))
}

func (h *Handler) UpdateOrder(c *fiber.Ctx) error {
	var req UpdateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return kernel.NewBadRequestError("invalid request body")
	}
	if err := sharedvalidator.ValidateStruct(req); err != nil {
		return err
	}

	result, err := h.useCase.UpdateOrder(c.UserContext(), c.Params("id"), toUpdateOrderDTO(req))
	if err != nil {
		return err
	}
	return c.JSON(toOrderResponse(result))
}

func (h *Handler) DeleteOrder(c *fiber.Ctx) error {
	if err := h.useCase.DeleteOrder(c.UserContext(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) GetOrder(c *fiber.Ctx) error {
	result, err := h.useCase.GetOrder(c.UserContext(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(toOrderResponse(result))
}

func (h *Handler) ListOrders(c *fiber.Ctx) error {
	result, err := h.useCase.ListOrders(c.UserContext())
	if err != nil {
		return err
	}
	items := make([]OrderResponse, 0, len(result))
	for _, order := range result {
		items = append(items, toOrderResponse(order))
	}
	return c.JSON(items)
}
