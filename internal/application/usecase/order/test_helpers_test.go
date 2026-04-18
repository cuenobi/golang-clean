package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/order"
	usecase "github.com/cuenobi/golang-clean/internal/application/usecase/order"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/event"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/stretchr/testify/require"
)

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

func (m *orderRepoMock) GetByIdempotencyKey(ctx context.Context, idempotencyKey string) (*entity.Order, error) {
	for _, order := range m.orders {
		if order.IdempotencyKey == idempotencyKey {
			return order, nil
		}
	}
	return nil, kernel.ErrNotFound
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

type outboxWriterMock struct {
	called int
}

func (m *outboxWriterMock) EnqueueOrderCreated(ctx context.Context, eventPayload event.OrderCreated) error {
	m.called++
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

func newOrderUseCaseForTest(id string, now time.Time) (*usecase.OrderUseCase, *orderRepoMock, *outboxWriterMock) {
	repo := newOrderRepoMock()
	outbox := &outboxWriterMock{}
	uc := usecase.NewOrderUseCase(repo, &txMock{}, outbox, fixedClock{now: now}, fixedID{id: id})
	return uc, repo, outbox
}

func seedOrder(
	t *testing.T,
	repo *orderRepoMock,
	id string,
	customerID string,
	currency string,
	amount int64,
	now time.Time,
) {
	t.Helper()

	money, err := valueobject.NewMoney(currency, amount)
	require.NoError(t, err)
	order, err := entity.NewOrder(id, customerID, "", money, now)
	require.NoError(t, err)
	require.NoError(t, repo.Save(context.Background(), order))
}

func createOrderReq(idempotencyKey string) dto.CreateOrderRequest {
	return dto.CreateOrderRequest{
		CustomerID:     "cus_1",
		Currency:       "USD",
		Amount:         100,
		IdempotencyKey: idempotencyKey,
	}
}
