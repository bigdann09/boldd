package services

import (
	"log"

	"github.com/boldd/internal/config"
	"github.com/boldd/internal/infrastructure/persistence"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewServices(cfg *config.Config) *Service {
	// register database
	db, err := persistence.NewDB(&cfg.DatabaseConfig)
	if err != nil {
		log.Println("could not connect to database")
		panic(err)
	}

	// register redis

	return &Service{
		DB: db,
	}
}
