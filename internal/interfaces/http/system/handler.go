package http

import (
	"context"
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	cfg config.Config
	db  *gorm.DB
}

func NewHandler(cfg config.Config, db *gorm.DB) *Handler {
	return &Handler{cfg: cfg, db: db}
}

func (h *Handler) Healthz(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": h.cfg.AppName,
		"env":     h.cfg.AppEnv,
	})
}

func (h *Handler) Readyz(c *fiber.Ctx) error {
	sqlDB, err := h.db.DB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status": "not_ready",
			"reason": "db_instance_unavailable",
		})
	}

	timeout := time.Duration(h.cfg.ReadinessDBTimeoutMS) * time.Millisecond
	if timeout <= 0 {
		timeout = 1500 * time.Millisecond
	}
	ctx, cancel := context.WithTimeout(c.UserContext(), timeout)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status": "not_ready",
			"reason": "db_ping_failed",
		})
	}

	return c.JSON(fiber.Map{
		"status": "ready",
	})
}
