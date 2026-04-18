package usecase

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
)

var _ in.OrderUseCase = (*OrderUseCase)(nil)

type OrderUseCase struct {
	repo   out.OrderRepository
	tx     out.TxManager
	outbox out.OrderEventOutboxWriter
	clock  out.Clock
	idGen  out.IDGenerator
}

func NewOrderUseCase(
	repo out.OrderRepository,
	tx out.TxManager,
	outbox out.OrderEventOutboxWriter,
	clock out.Clock,
	idGen out.IDGenerator,
) *OrderUseCase {
	return &OrderUseCase{
		repo:   repo,
		tx:     tx,
		outbox: outbox,
		clock:  clock,
		idGen:  idGen,
	}
}
