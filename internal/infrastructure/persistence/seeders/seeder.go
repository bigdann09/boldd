package seeder

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Seeder struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewSeeder(db *gorm.DB, logger *zap.Logger) *Seeder {
	return &Seeder{db, logger}
}

func (seeder *Seeder) Run() {
	seeder.roleSeeder()
}
