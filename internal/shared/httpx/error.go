package httpx

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		status := fiber.StatusInternalServerError
		code := "internal_error"

		if e, ok := err.(*fiber.Error); ok {
			status = e.Code
			switch {
			case e.Code >= 500:
				code = "internal_error"
			case e.Code == fiber.StatusNotFound:
				code = "not_found"
			case e.Code == fiber.StatusConflict:
				code = "conflict"
			case e.Code == fiber.StatusBadRequest:
				code = "bad_request"
			default:
				code = "request_error"
			}
		} else {
			status = fiber.StatusBadRequest
			code = "bad_request"
		}

		return c.Status(status).JSON(ErrorResponse{
			Code:    code,
			Message: err.Error(),
		})
	}
}
