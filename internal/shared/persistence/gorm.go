package persistence

import (
	"fmt"
	"strings"
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

	tz := cfg.PostgresTZ
	if strings.TrimSpace(tz) == "" {
		tz = "UTC"
	}
	tz = strings.ReplaceAll(tz, "'", "''")

	if err := db.Exec(fmt.Sprintf("SET TIME ZONE '%s'", tz)).Error; err != nil {
		return nil, err
	}

	return db, nil
}
