package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestListOrders(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, repo, _ := newOrderUseCaseForTest("ord_123", now)

	_, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)
	seedOrder(t, repo, "ord_456", "cus_2", "THB", 500, now)

	items, err := uc.ListOrders(context.Background())
	require.NoError(t, err)
	require.Len(t, items, 2)
}
