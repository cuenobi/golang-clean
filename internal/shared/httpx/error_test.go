package httpx

import (
	"testing"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/gofiber/fiber/v2"
)

func TestMapToAppError_MapsFiber413(t *testing.T) {
	appErr := mapToAppError(fiber.NewError(fiber.StatusRequestEntityTooLarge, "body too large"))
	if appErr.Code != kernel.ErrorCodePayloadTooLarge {
		t.Fatalf("expected code %d, got %d", kernel.ErrorCodePayloadTooLarge, appErr.Code)
	}
	if appErr.HTTPStatus != fiber.StatusRequestEntityTooLarge {
		t.Fatalf("expected status %d, got %d", fiber.StatusRequestEntityTooLarge, appErr.HTTPStatus)
	}
}

func TestMapToAppError_MapsFiber415(t *testing.T) {
	appErr := mapToAppError(fiber.NewError(fiber.StatusUnsupportedMediaType, "unsupported media type"))
	if appErr.Code != kernel.ErrorCodeUnsupportedMediaType {
		t.Fatalf("expected code %d, got %d", kernel.ErrorCodeUnsupportedMediaType, appErr.Code)
	}
	if appErr.HTTPStatus != fiber.StatusUnsupportedMediaType {
		t.Fatalf("expected status %d, got %d", fiber.StatusUnsupportedMediaType, appErr.HTTPStatus)
	}
}
