package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDeleteOrder(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _, _ := newOrderUseCaseForTest("ord_123", now)

	_, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)

	err = uc.DeleteOrder(context.Background(), "ord_123")
	require.NoError(t, err)

	_, err = uc.GetOrder(context.Background(), "ord_123")
	require.Error(t, err)
}
