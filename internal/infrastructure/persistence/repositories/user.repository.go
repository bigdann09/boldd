package repositories

import (
	"fmt"

	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *entities.User) error
	Find(id int) (interface{}, error)
	EmailExists(email string) bool
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

func (repo UserRepository) Find(id int) (interface{}, error) {
	var response interface{}
	result := repo.db.Table("users").Where("id = ?", id).Scan(&response)
	return response, result.Error
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
