package http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	group := router.Group("/users")
	group.Post("", handler.CreateUser)
	group.Get("", handler.ListUsers)
	group.Get("/:id", handler.GetUser)
	group.Put("/:id", handler.UpdateUser)
	group.Delete("/:id", handler.DeleteUser)
}
