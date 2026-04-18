package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _ := newUserUseCaseForTest("usr_1", now)

	created, err := uc.CreateUser(context.Background(), createUserReq())
	require.NoError(t, err)
	require.Equal(t, "usr_1", created.ID)
}
