package repositories

import (
	"github.com/boldd/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) Create(user *user.User) error {
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
