package usecase_test

import (
	"context"
	"testing"
	"time"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _ := newUserUseCaseForTest("usr_1", now)

	_, err := uc.CreateUser(context.Background(), createUserReq())
	require.NoError(t, err)

	updated, err := uc.UpdateUser(context.Background(), "usr_1", dto.UpdateUserRequest{
		Name:  "Alice Updated",
		Email: "alice.updated@example.com",
	})
	require.NoError(t, err)
	require.Equal(t, "Alice Updated", updated.Name)
	require.Equal(t, "alice.updated@example.com", updated.Email)
}
