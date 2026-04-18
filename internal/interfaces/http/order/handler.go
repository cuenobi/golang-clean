package http

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	sharedhttpx "github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedvalidator "github.com/cuenobi/golang-clean/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase in.OrderUseCase
}

type ErrorResponseDoc = sharedhttpx.ErrorResponse

func NewHandler(useCase in.OrderUseCase) *Handler {
	return &Handler{useCase: useCase}
}

// CreateOrder godoc
// @Summary Create order
// @Description Create a new order (idempotent when Idempotency-Key is provided).
// @Tags Orders
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param X-Permissions header string true "Required permission: orders:write"
// @Param Idempotency-Key header string false "Idempotency key"
// @Param request body CreateOrderRequest true "Create order payload"
// @Success 201 {object} OrderResponse
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 409 {object} ErrorResponseDoc
// @Failure 422 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/orders [post]
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

// UpdateOrder godoc
// @Summary Update order
// @Tags Orders
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param X-Permissions header string true "Required permission: orders:write"
// @Param id path string true "Order ID"
// @Param request body UpdateOrderRequest true "Update order payload"
// @Success 200 {object} OrderResponse
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 422 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/orders/{id} [put]
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

// DeleteOrder godoc
// @Summary Delete order
// @Tags Orders
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: orders:write"
// @Param id path string true "Order ID"
// @Success 204
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/orders/{id} [delete]
func (h *Handler) DeleteOrder(c *fiber.Ctx) error {
	if err := h.useCase.DeleteOrder(c.UserContext(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// GetOrder godoc
// @Summary Get order by ID
// @Tags Orders
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: orders:read"
// @Param id path string true "Order ID"
// @Success 200 {object} OrderResponse
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/orders/{id} [get]
func (h *Handler) GetOrder(c *fiber.Ctx) error {
	result, err := h.useCase.GetOrder(c.UserContext(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(toOrderResponse(result))
}

// ListOrders godoc
// @Summary List orders
// @Tags Orders
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: orders:read"
// @Success 200 {array} OrderResponse
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/orders [get]
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
