package persistence

import (
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
