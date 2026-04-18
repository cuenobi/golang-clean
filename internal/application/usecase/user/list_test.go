package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, repo := newUserUseCaseForTest("usr_1", now)

	_, err := uc.CreateUser(context.Background(), createUserReq())
	require.NoError(t, err)
	seedUser(t, repo, "usr_2", "Bob", "bob@example.com", now)

	items, err := uc.ListUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, items, 2)
}
