package di

import (
	"time"

	_ "github.com/cuenobi/golang-clean/api/swagger"
	httpadapter "github.com/cuenobi/golang-clean/internal/interfaces/http/order"
	systemhttp "github.com/cuenobi/golang-clean/internal/interfaces/http/system"
	userhttp "github.com/cuenobi/golang-clean/internal/interfaces/http/user"
	"github.com/cuenobi/golang-clean/internal/shared/httpx"
	"github.com/cuenobi/golang-clean/internal/shared/metrics"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
)

func (c *Container) wireHTTP() {
	readTimeout := time.Duration(c.Cfg.HTTPReadTimeout) * time.Second
	writeTimeout := time.Duration(c.Cfg.HTTPWriteTimeout) * time.Second
	bodyLimit := c.Cfg.HTTPBodyLimitMB * 1024 * 1024
	if bodyLimit <= 0 {
		bodyLimit = 4 * 1024 * 1024
	}

	app := fiber.New(fiber.Config{
		ErrorHandler:  httpx.NewErrorHandler(),
		ReadTimeout:   readTimeout,
		WriteTimeout:  writeTimeout,
		BodyLimit:     bodyLimit,
		CaseSensitive: true,
	})

	app.Use(httpx.RequestIDMiddleware())
	app.Use(httpx.CORSMiddleware(c.Cfg))
	app.Use(metrics.HTTPMiddleware())
	app.Use(httpx.RateLimiter(c.Cfg))
	app.Use(httpx.RequestLogger(c.Log))

	systemHandler := systemhttp.NewHandler(c.Cfg, c.DB)
	systemhttp.RegisterRoutes(app, systemHandler)
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	v1 := app.Group("/api/v1")
	httpHandler := httpadapter.NewHandler(c.OrderUseCase)
	httpadapter.RegisterRoutes(v1, httpHandler, c.Cfg)
	userHandler := userhttp.NewHandler(c.UserUseCase)
	userhttp.RegisterRoutes(v1, userHandler, c.Cfg)

	c.HTTPApp = app
}
