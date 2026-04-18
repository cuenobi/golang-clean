package http

import (
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler, cfg config.Config) {
	group := router.Group("/users", httpx.APIKeyAuth(cfg))
	group.Post("", httpx.RequirePermission(cfg, "users:write"), handler.CreateUser)
	group.Get("", httpx.RequirePermission(cfg, "users:read"), handler.ListUsers)
	group.Get("/:id", httpx.RequirePermission(cfg, "users:read"), handler.GetUser)
	group.Put("/:id", httpx.RequirePermission(cfg, "users:write"), handler.UpdateUser)
	group.Delete("/:id", httpx.RequirePermission(cfg, "users:write"), handler.DeleteUser)
}
