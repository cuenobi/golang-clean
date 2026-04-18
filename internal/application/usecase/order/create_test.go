package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _, outbox := newOrderUseCaseForTest("ord_123", now)

	created, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)
	require.Equal(t, "ord_123", created.ID)
	require.Equal(t, 1, outbox.called)

	duplicate, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)
	require.Equal(t, created.ID, duplicate.ID)
	require.Equal(t, 1, outbox.called)
}
