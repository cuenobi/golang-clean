package httpx

import (
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func RequestIDMiddleware() fiber.Handler {
	return requestid.New(requestid.Config{
		Header: "X-Request-ID",
	})
}

func RequestLogger(log logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		status := c.Response().StatusCode()

		fields := map[string]any{
			"request_id": c.GetRespHeader("X-Request-ID"),
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     status,
			"latency_ms": time.Since(start).Milliseconds(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}

		switch {
		case status >= fiber.StatusInternalServerError:
			log.Error("http_request_failed", err, fields)
		case status >= fiber.StatusBadRequest:
			log.Warn("http_request_client_error", fields)
		default:
			log.Info("http_request", fields)
		}

		return err
	}
}
