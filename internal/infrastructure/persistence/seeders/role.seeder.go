package seeder

import (
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"go.uber.org/zap"
)

func (seeder *Seeder) roleSeeder() {
	roleRepository := repositories.NewRoleRepository(seeder.db)
	seeder.logger.Info("counting total number of roles in database")
	count, _ := roleRepository.Count()
	if count == 0 {
		seeder.logger.Info("seeding to database since count is zero(0)")
		roles := []string{"customer", "admin", "vendor"}
		for _, name := range roles {
			seeder.logger.Info("adding role to database", zap.String("role", name))
			roleRepository.Create(entities.NewRole(name))
		}
	}
	seeder.logger.Info("skipping seeding for roles")
}
