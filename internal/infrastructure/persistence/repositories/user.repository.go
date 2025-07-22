package repositories

import (
	"errors"
	"fmt"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *entities.User) error
	Find(id int) (dtos.UserResponse, error)
	FindByEmail(email string) (dtos.UserResponse, error)
	Update(id uint, user *entities.User) error
	EmailExists(email string) bool
	Roles(userID uint) ([]string, error)
	Delete(id int) error
	AssignRole(userID int, role string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) Create(user *entities.User) error {
	result := repo.db.Table("users").Create(&user)
	return result.Error
}

func (repo UserRepository) Find(id int) (dtos.UserResponse, error) {
	var data dtos.User
	var response dtos.UserResponse
	result := repo.db.Table("users").Where("id = ?", id).Scan(&data)
	if result.Error == nil {
		roles, _ := repo.Roles(data.ID)
		response = data.Convert(roles...)
	}
	return response, result.Error
}

func (repo UserRepository) FindByEmail(email string) (dtos.UserResponse, error) {
	var data dtos.User
	var response dtos.UserResponse
	result := repo.db.Table("users").Where("email = ?", email).Scan(&data)
	if result.Error == nil {
		roles, _ := repo.Roles(data.ID)
		response = data.Convert(roles...)
	}
	return response, result.Error
}

func (repo UserRepository) Update(id uint, user *entities.User) error {
	result := repo.db.Table("users").Where("id = ?", id).Updates(user)
	return result.Error
}

func (repo UserRepository) EmailExists(email string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from users where email = ?)", email).Scan(&exists)
	return exists
}

func (repo UserRepository) AssignRole(userID int, role string) error {
	// check if role exists and retrieve data
	roleRepository := NewRoleRepository(repo.db)
	data, err := roleRepository.FindByName(role)
	if err != nil {
		return err
	}

	if data.Name == "" || data.ID == 0 {
		return fmt.Errorf("role %s does not exists", role)
	}

	result := repo.db.Exec("INSERT INTO user_roles(user_id, role_id) VALUES(?, ?)", userID, data.ID)
	return result.Error
}

func (repo UserRepository) Roles(userID uint) ([]string, error) {
	var roles []string
	result := repo.db.Raw("select roles.name from users left join user_roles ON users.id = user_roles.user_id left join roles on user_roles.role_id = roles.id where users.id = ?", userID).Scan(&roles)
	return roles, result.Error
}

func (repo UserRepository) Delete(id int) error {
	result := repo.db.Table("users").Unscoped().Where("id = ?", id).Delete(&entities.User{})
	if result.RowsAffected == 0 {
		return errors.New("record not deleted")
	}
	return result.Error
}
