package httpx

import (
	"crypto/subtle"
	"strings"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/gofiber/fiber/v2"
)

const (
	headerAPIKey      = "X-API-Key"
	headerPermissions = "X-Permissions"
)

func APIKeyAuth(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !cfg.AuthEnabled {
			return c.Next()
		}
		if cfg.APIKey == "" {
			return kernel.NewInternalError()
		}

		provided := c.Get(headerAPIKey)
		if provided == "" {
			return kernel.NewUnauthorizedError("missing api key")
		}
		if subtle.ConstantTimeCompare([]byte(provided), []byte(cfg.APIKey)) != 1 {
			return kernel.NewUnauthorizedError("invalid api key")
		}
		return c.Next()
	}
}

func RequirePermission(cfg config.Config, required string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !cfg.AuthEnabled {
			return c.Next()
		}
		raw := c.Get(headerPermissions)
		if raw == "" {
			return kernel.NewForbiddenError("missing permissions")
		}

		perms := strings.Split(raw, ",")
		for _, permission := range perms {
			if strings.TrimSpace(permission) == required {
				return c.Next()
			}
		}

		return kernel.NewForbiddenError("permission denied")
	}
}
