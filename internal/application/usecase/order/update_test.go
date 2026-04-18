package usecase_test

import (
	"context"
	"testing"
	"time"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
	"github.com/stretchr/testify/require"
)

func TestUpdateOrder(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	uc, _, _ := newOrderUseCaseForTest("ord_123", now)

	_, err := uc.CreateOrder(context.Background(), createOrderReq("idem-1"))
	require.NoError(t, err)

	updated, err := uc.UpdateOrder(context.Background(), "ord_123", dto.UpdateOrderRequest{
		CustomerID: "cus_2",
		Currency:   "THB",
		Amount:     200,
	})
	require.NoError(t, err)
	require.Equal(t, "cus_2", updated.CustomerID)
	require.Equal(t, "THB", updated.Currency)
	require.Equal(t, int64(200), updated.Amount)
}
