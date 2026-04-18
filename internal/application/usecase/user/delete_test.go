package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDeleteUser(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _ := newUserUseCaseForTest("usr_1", now)

	_, err := uc.CreateUser(context.Background(), createUserReq())
	require.NoError(t, err)

	err = uc.DeleteUser(context.Background(), "usr_1")
	require.NoError(t, err)

	_, err = uc.GetUser(context.Background(), "usr_1")
	require.Error(t, err)
}
