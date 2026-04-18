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

		fields := map[string]any{
			"request_id": c.GetRespHeader("X-Request-ID"),
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency_ms": time.Since(start).Milliseconds(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}

		if err != nil {
			log.Error("http_request_failed", err, fields)
			return err
		}

		log.Info("http_request", fields)
		return nil
	}
}
