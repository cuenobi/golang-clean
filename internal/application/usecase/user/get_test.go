package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _ := newUserUseCaseForTest("usr_1", now)

	_, err := uc.CreateUser(context.Background(), createUserReq())
	require.NoError(t, err)

	got, err := uc.GetUser(context.Background(), "usr_1")
	require.NoError(t, err)
	require.Equal(t, "Alice", got.Name)
}
