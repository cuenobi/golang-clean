package di

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"gorm.io/gorm"
)

type Module struct {
	HTTPRunner     func() error
	ConsumerRunner func(context.Context) error
}

func NewModule(cfg config.Config, db *gorm.DB) (*Module, error) {
	container, err := NewContainer(cfg, db)
	if err != nil {
		return nil, err
	}

	return &Module{
		HTTPRunner: func() error {
			defer container.Close()
			return container.HTTPApp.Listen(cfg.HTTPAddress)
		},
		ConsumerRunner: func(ctx context.Context) error {
			defer container.Close()
			return container.Consumer.Run(ctx)
		},
	}, nil
}
