package repositories

import (
	"errors"

	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IOtpRepository interface {
	Create(otp *entities.Otp) error
	Find(email string) (entities.Otp, error)
	Delete(id int) error
}
type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) *OtpRepository {
	return &OtpRepository{db}
}

func (repo OtpRepository) Create(otp *entities.Otp) error {
	result := repo.db.Table("otp").Create(&otp)
	return result.Error
}

func (repo OtpRepository) Find(email string) (entities.Otp, error) {
	var response entities.Otp
	result := repo.db.Table("otp").Where("email = ?", email).Scan(&response)
	return response, result.Error
}

func (repo OtpRepository) Delete(id int) error {
	result := repo.db.Table("otp").Unscoped().Where("id = ?", id).Delete(&entities.Otp{})
	if result.RowsAffected == 0 {
		return errors.New("record not deleted")
	}
	return result.Error
}
