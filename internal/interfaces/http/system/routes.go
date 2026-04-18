package http

import (
	"github.com/cuenobi/golang-clean/internal/shared/metrics"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	router.Get("/healthz", handler.Healthz)
	router.Get("/readyz", handler.Readyz)
	router.Get("/metrics", metrics.Handler())
}
