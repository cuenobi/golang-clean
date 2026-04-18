package validator

import (
	"testing"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
)

type samplePayload struct {
	Email string `json:"email" validate:"required,email"`
}

func TestValidateStruct_ReturnsValidationData(t *testing.T) {
	err := ValidateStruct(samplePayload{Email: "invalid-email"})
	if err == nil {
		t.Fatal("expected validation error")
	}

	appErr, ok := err.(*kernel.AppError)
	if !ok {
		t.Fatalf("expected *kernel.AppError, got %T", err)
	}
	if appErr.Code != kernel.ErrorCodeValidation {
		t.Fatalf("expected code %d, got %d", kernel.ErrorCodeValidation, appErr.Code)
	}
	if appErr.Data == nil {
		t.Fatal("expected validation data")
	}

	data, ok := appErr.Data.(map[string]any)
	if !ok {
		t.Fatalf("expected map data, got %T", appErr.Data)
	}
	violations, ok := data["violations"].([]FieldViolation)
	if !ok {
		t.Fatalf("expected []FieldViolation, got %T", data["violations"])
	}
	if len(violations) == 0 {
		t.Fatal("expected at least one violation")
	}
	if violations[0].Field != "email" {
		t.Fatalf("expected field email, got %s", violations[0].Field)
	}
}
