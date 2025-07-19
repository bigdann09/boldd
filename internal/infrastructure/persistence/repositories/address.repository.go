package repositories

import (
	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IUserAddressRepository interface {
	Create(address *entities.UserAddress) error
}

type UserAddressRepository struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) *UserAddressRepository {
	return &UserAddressRepository{db}
}

func (repo UserAddressRepository) Create(address *entities.UserAddress) error {
	result := repo.db.Table("user_addresses").Create(&address)
	return result.Error
}
