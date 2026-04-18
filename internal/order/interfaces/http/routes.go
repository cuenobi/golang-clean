package http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	group := router.Group("/orders")
	group.Post("", handler.CreateOrder)
	group.Get("", handler.ListOrders)
	group.Get("/:id", handler.GetOrder)
	group.Put("/:id", handler.UpdateOrder)
	group.Delete("/:id", handler.DeleteOrder)
}
