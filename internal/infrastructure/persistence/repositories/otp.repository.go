package repositories

import (
	"errors"

	"github.com/boldd/internal/domain/entities"
	"gorm.io/gorm"
)

type IOtpRepository interface {
	Create(otp *entities.Otp) error
	Find(email string) (entities.Otp, error)
	Exists(email string) bool
	Delete(uuid string) error
	DeleteByEmail(email string) error
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

func (repo OtpRepository) Exists(email string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from users where email = ?)", email).Scan(&exists)
	return exists
}

func (repo OtpRepository) Delete(uuid string) error {
	result := repo.db.Table("otp").Unscoped().Where("uuid = ?", uuid).Delete(&entities.Otp{})
	if result.RowsAffected == 0 {
		return errors.New("record not deleted")
	}
	return result.Error
}

func (repo OtpRepository) DeleteByEmail(email string) error {
	result := repo.db.Table("otp").Unscoped().Where("email = ?", email).Delete(&entities.Otp{})
	if result.RowsAffected == 0 {
		return errors.New("record not deleted")
	}
	return result.Error
}
