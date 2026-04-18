package persistence

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"gorm.io/gorm"
)

type gormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) kernel.TxManager {
	return &gormTxManager{db: db}
}

func (m *gormTxManager) WithinTransaction(ctx context.Context, fn func(context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := WithTx(ctx, tx)
		return fn(txCtx)
	})
}
