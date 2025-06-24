package services

import (
	"log"

	"github.com/boldd/internal/config"
	"github.com/boldd/internal/infrastructure/persistence"
	"github.com/boldd/pkgs/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewServices(cfg *config.Config) *Service {
	// register database
	db, err := persistence.NewDB(&cfg.DatabaseConfig)
	if err != nil {
		log.Println("could not connect to database")
		panic(err)
	}

	// register logger
	logger := logger.NewLogger(cfg.Environment)
	defer logger.Sync()

	// register redis

	return &Service{
		DB:     db,
		Logger: logger,
	}
}
