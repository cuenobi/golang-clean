package http

import (
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler, cfg config.Config) {
	group := router.Group("/audit-logs", httpx.APIKeyAuth(cfg))
	group.Get("/system", httpx.RequirePermission(cfg, "audit_logs:read"), handler.GetSystemAuditLogs)
	group.Get("/organization", httpx.RequirePermission(cfg, "audit_logs:read"), handler.GetOrganizationAuditLogs)
}
