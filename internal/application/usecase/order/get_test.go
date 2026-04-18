package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetOrder(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _, _ := newOrderUseCaseForTest("ord_123", now)

	_, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)

	got, err := uc.GetOrder(context.Background(), "ord_123")
	require.NoError(t, err)
	require.Equal(t, "cus_1", got.CustomerID)
}
