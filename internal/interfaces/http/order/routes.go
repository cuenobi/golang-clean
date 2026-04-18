package http

import (
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler, cfg config.Config) {
	group := router.Group("/orders", httpx.APIKeyAuth(cfg))
	group.Post("", httpx.RequirePermission(cfg, "orders:write"), handler.CreateOrder)
	group.Get("", httpx.RequirePermission(cfg, "orders:read"), handler.ListOrders)
	group.Get("/:id", httpx.RequirePermission(cfg, "orders:read"), handler.GetOrder)
	group.Put("/:id", httpx.RequirePermission(cfg, "orders:write"), handler.UpdateOrder)
	group.Delete("/:id", httpx.RequirePermission(cfg, "orders:write"), handler.DeleteOrder)
}
