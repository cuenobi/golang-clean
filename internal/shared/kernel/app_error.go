package kernel

import "github.com/gofiber/fiber/v2"

type AppError struct {
	HTTPStatus int
	Code       int
	Type       string
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Type
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewValidationError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusBadRequest, Code: ErrorCodeValidation, Type: "validation_error", Message: message}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusBadRequest, Code: ErrorCodeBadRequest, Type: "bad_request", Message: message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusNotFound, Code: ErrorCodeNotFound, Type: "not_found", Message: message}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusUnauthorized, Code: ErrorCodeUnauthorized, Type: "unauthorized", Message: message}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusForbidden, Code: ErrorCodeForbidden, Type: "forbidden", Message: message}
}

func NewConflictError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusConflict, Code: ErrorCodeConflict, Type: "conflict", Message: message}
}

func NewInvalidStateError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusUnprocessableEntity, Code: ErrorCodeInvalidState, Type: "invalid_state", Message: message}
}

func NewRateLimitError(message string) *AppError {
	return &AppError{HTTPStatus: fiber.StatusTooManyRequests, Code: ErrorCodeRateLimited, Type: "rate_limited", Message: message}
}

func NewInternalError() *AppError {
	return &AppError{HTTPStatus: fiber.StatusInternalServerError, Code: ErrorCodeInternal, Type: "internal_error", Message: "internal server error"}
}
