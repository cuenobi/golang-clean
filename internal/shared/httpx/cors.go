package httpx

import (
	"strings"
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware(cfg config.Config) fiber.Handler {
	origins := cfg.CORSOrigins
	if len(origins) == 0 {
		origins = []string{
			"http://localhost:3000",
			"http://localhost:5173",
		}
	}

	methods := cfg.CORSMethods
	if len(methods) == 0 {
		methods = []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		}
	}

	headers := cfg.CORSHeaders
	if len(headers) == 0 {
		headers = []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-API-Key",
			"X-Request-ID",
			"Idempotency-Key",
		}
	}

	expose := cfg.CORSExpose
	if len(expose) == 0 {
		expose = []string{"X-Request-ID"}
	}

	maxAge := time.Duration(cfg.CORSMaxAgeSec) * time.Second
	if maxAge < 0 {
		maxAge = 0
	}

	return cors.New(cors.Config{
		AllowOrigins:     strings.Join(origins, ","),
		AllowMethods:     strings.Join(methods, ","),
		AllowHeaders:     strings.Join(headers, ","),
		ExposeHeaders:    strings.Join(expose, ","),
		AllowCredentials: cfg.CORSAllowCreds,
		MaxAge:           int(maxAge.Seconds()),
	})
}
