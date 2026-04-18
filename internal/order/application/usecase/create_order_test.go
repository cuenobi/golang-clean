package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cuenobi/golang-clean/internal/order/application/dto"
	"github.com/cuenobi/golang-clean/internal/order/application/usecase"
	"github.com/cuenobi/golang-clean/internal/order/domain/entity"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type orderUseCaseSuite struct {
	suite.Suite
}

func TestOrderUseCaseSuite(t *testing.T) {
	suite.Run(t, new(orderUseCaseSuite))
}

func (s *orderUseCaseSuite) TestCRUDFlow() {
	repo := newOrderRepoMock()
	tx := &txMock{}
	publisher := &publisherMock{}
	clock := fixedClock{now: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	idGen := fixedID{id: "ord_123"}

	uc := usecase.NewOrderUseCase(repo, tx, publisher, clock, idGen)

	created, err := uc.CreateOrder(context.Background(), dto.CreateOrderRequest{
		CustomerID: "cus_1",
		Currency:   "USD",
		Amount:     100,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), "ord_123", created.ID)
	require.True(s.T(), publisher.called)

	got, err := uc.GetOrder(context.Background(), "ord_123")
	require.NoError(s.T(), err)
	require.Equal(s.T(), "cus_1", got.CustomerID)

	updated, err := uc.UpdateOrder(context.Background(), "ord_123", dto.UpdateOrderRequest{
		CustomerID: "cus_2",
		Currency:   "THB",
		Amount:     200,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), "cus_2", updated.CustomerID)
	require.Equal(s.T(), "THB", updated.Currency)
	require.Equal(s.T(), int64(200), updated.Amount)

	list, err := uc.ListOrders(context.Background())
	require.NoError(s.T(), err)
	require.Len(s.T(), list, 1)

	err = uc.DeleteOrder(context.Background(), "ord_123")
	require.NoError(s.T(), err)

	_, err = uc.GetOrder(context.Background(), "ord_123")
	require.Error(s.T(), err)
}

// Intentionally small hand-rolled mocks for bootstrap tests.
type orderRepoMock struct {
	orders map[string]*entity.Order
}

func newOrderRepoMock() *orderRepoMock {
	return &orderRepoMock{orders: map[string]*entity.Order{}}
}

func (m *orderRepoMock) Save(ctx context.Context, order *entity.Order) error {
	m.orders[order.ID] = order
	return nil
}

func (m *orderRepoMock) GetByID(ctx context.Context, orderID string) (*entity.Order, error) {
	order, ok := m.orders[orderID]
	if !ok {
		return nil, kernel.ErrNotFound
	}
	return order, nil
}

func (m *orderRepoMock) List(ctx context.Context) ([]*entity.Order, error) {
	result := make([]*entity.Order, 0, len(m.orders))
	for _, order := range m.orders {
		result = append(result, order)
	}
	return result, nil
}

func (m *orderRepoMock) Update(ctx context.Context, order *entity.Order) error {
	if _, ok := m.orders[order.ID]; !ok {
		return errors.New("not found")
	}
	m.orders[order.ID] = order
	return nil
}

func (m *orderRepoMock) Delete(ctx context.Context, orderID string) error {
	if _, ok := m.orders[orderID]; !ok {
		return kernel.ErrNotFound
	}
	delete(m.orders, orderID)
	return nil
}

type txMock struct{}

func (m *txMock) WithinTransaction(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

type publisherMock struct {
	called bool
}

func (m *publisherMock) PublishOrderCreated(ctx context.Context, eventPayload any) error {
	m.called = true
	return nil
}

type fixedClock struct {
	now time.Time
}

func (f fixedClock) Now() time.Time {
	return f.now
}

type fixedID struct {
	id string
}

func (f fixedID) NewID() string {
	return f.id
}
