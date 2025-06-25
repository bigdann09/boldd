package persistence

import (
	"fmt"

	"github.com/boldd/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(formatToDSN(cfg)))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func formatToDSN(cfg *config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Port,
		cfg.SSLMode,
	)
}
