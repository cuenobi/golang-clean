package persistence

import (
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN()), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	if err := db.Exec("SET TIME ZONE ?", cfg.PostgresTZ).Error; err != nil {
		return nil, err
	}

	return db, nil
}
