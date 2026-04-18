package persistence

import (
	"context"

	"gorm.io/gorm"
)

type txContextKey struct{}

func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txContextKey{}, tx)
}

func FromContext(ctx context.Context, fallback *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txContextKey{}).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return fallback
}
