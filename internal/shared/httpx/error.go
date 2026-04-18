package httpx

import (
	"errors"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code      int    `json:"code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		appErr := mapToAppError(err)
		requestID := c.GetRespHeader("X-Request-ID")

		return c.Status(appErr.HTTPStatus).JSON(ErrorResponse{
			Code:      appErr.Code,
			Type:      appErr.Type,
			Message:   appErr.Message,
			Data:      appErr.Data,
			RequestID: requestID,
		})
	}
}

func mapToAppError(err error) *kernel.AppError {
	var appErr *kernel.AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	switch {
	case errors.Is(err, kernel.ErrNotFound):
		return kernel.NewNotFoundError("resource not found")
	case errors.Is(err, kernel.ErrConflict):
		return kernel.NewConflictError("resource conflict")
	case errors.Is(err, kernel.ErrInvalidState):
		return kernel.NewInvalidStateError("invalid state")
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		switch fiberErr.Code {
		case fiber.StatusBadRequest:
			return kernel.NewBadRequestError(fiberErr.Message)
		case fiber.StatusRequestEntityTooLarge:
			return kernel.NewPayloadTooLargeError(fiberErr.Message)
		case fiber.StatusUnsupportedMediaType:
			return kernel.NewUnsupportedMediaTypeError(fiberErr.Message)
		case fiber.StatusUnauthorized:
			return kernel.NewUnauthorizedError(fiberErr.Message)
		case fiber.StatusForbidden:
			return kernel.NewForbiddenError(fiberErr.Message)
		case fiber.StatusNotFound:
			return kernel.NewNotFoundError(fiberErr.Message)
		case fiber.StatusConflict:
			return kernel.NewConflictError(fiberErr.Message)
		case fiber.StatusTooManyRequests:
			return kernel.NewRateLimitError(fiberErr.Message)
		default:
			if fiberErr.Code >= 500 {
				return kernel.NewInternalError()
			}
			return &kernel.AppError{
				HTTPStatus: fiberErr.Code,
				Code:       kernel.ErrorCodeBadRequest,
				Type:       "request_error",
				Message:    fiberErr.Message,
			}
		}
	}

	return kernel.NewInternalError()
}
