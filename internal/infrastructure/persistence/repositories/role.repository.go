package repositories

import (
	"errors"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	Count() (int, error)
	Create(role *entities.Role) error
	RoleExists(name string) bool
	Find(id int) (interface{}, error)
	Update(uuid string, role *entities.Role) error
	FindByName(name string) (dtos.RoleResponse, error)
}
type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (repo RoleRepository) Create(role *entities.Role) error {
	result := repo.db.Table("roles").Create(&role)
	return result.Error
}

func (repo RoleRepository) Find(id int) (dtos.RoleResponse, error) {
	var response dtos.RoleResponse
	result := repo.db.Table("roles").Where("id = ?", id).Scan(&response)
	return response, result.Error
}

func (repo RoleRepository) FindByName(name string) (dtos.RoleResponse, error) {
	var response dtos.RoleResponse
	result := repo.db.Table("roles").Where("name = ?", name).Scan(&response)
	return response, result.Error
}

func (repo RoleRepository) RoleExists(name string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from roles where name = ?)", name).Scan(&exists)
	return exists
}

func (repo RoleRepository) Update(uuid string, role *entities.Role) error {
	result := repo.db.Table("roles").Where("uuid = ?", uuid).Updates(&role)
	if result.RowsAffected == 0 {
		return errors.New("role name not updated")
	}
	return result.Error
}

func (repo RoleRepository) Count() (int, error) {
	var count int64
	result := repo.db.Table("roles").Count(&count)
	return int(count), result.Error
}
