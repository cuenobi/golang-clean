package httpx

import (
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(cfg config.Config) fiber.Handler {
	if !cfg.RateLimitEnabled {
		return func(c *fiber.Ctx) error { return c.Next() }
	}

	window := time.Duration(cfg.RateLimitWindow) * time.Second
	if window <= 0 {
		window = 60 * time.Second
	}
	max := cfg.RateLimitMax
	if max <= 0 {
		max = 120
	}

	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: window,
		Next: func(c *fiber.Ctx) bool {
			path := c.Path()
			return path == "/healthz" || path == "/readyz" || path == "/metrics"
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			if key := c.Get(headerAPIKey); key != "" {
				return key
			}
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return kernel.NewRateLimitError("rate limit exceeded")
		},
	})
}
