package bootstrap

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/order/infrastructure/di"
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/persistence"
	"gorm.io/gorm"
)

type App struct {
	Config     config.Config
	DB         *gorm.DB
	HTTPRunner func() error
	Consumer   func(context.Context) error
}

func NewApp() (*App, error) {
	cfg := config.Load()
	db, err := persistence.NewGormDB(cfg)
	if err != nil {
		return nil, err
	}

	module, err := di.NewModule(cfg, db)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:     cfg,
		DB:         db,
		HTTPRunner: module.HTTPRunner,
		Consumer:   module.ConsumerRunner,
	}, nil
}

func (a *App) RunHTTP() error {
	return a.HTTPRunner()
}

func (a *App) RunConsumer(ctx context.Context) error {
	return a.Consumer(ctx)
}

func (a *App) Close(ctx context.Context) error {
	sqlDB, err := a.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
