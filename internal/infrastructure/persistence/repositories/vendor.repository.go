package repositories

import (
	"fmt"
	"strings"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/pkgs/utils"
	"gorm.io/gorm"
)

type IVendorRepository interface {
	Delete(id string) error
	VendorExists(name string) bool
	VendorExistsByID(id string) bool
	Create(address *entities.Vendor) error
	Find(id string) (*dtos.VendorResponse, error)
	Update(id string, Vendor *entities.Vendor) error
	FindAllPaginated(filter *dtos.VendorQueryFilter) (utils.PaginationResponse[dtos.VendorResponse], error)
}

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{db}
}

func (repo VendorRepository) Create(address *entities.Vendor) error {
	result := repo.db.Table("vendors").Create(&address)
	return result.Error
}

func (repo VendorRepository) FindAllPaginated(filter *dtos.VendorQueryFilter) (utils.PaginationResponse[dtos.VendorResponse], error) {
	if strings.EqualFold(filter.SortBy, "") {
		filter.SortBy = "name"
	}

	if strings.EqualFold(filter.Order, "") {
		filter.Order = "asc"
	}

	query := repo.db.Table("vendors").Order(fmt.Sprintf("%s %s", filter.SortBy, filter.Order))
	return utils.NewPaginationResponse[dtos.VendorResponse](filter.Page, filter.PageSize, query)
}

func (repo VendorRepository) Find(id string) (*dtos.VendorResponse, error) {
	var response *dtos.VendorResponse
	result := repo.db.Table("vendors").Where("id = ?", id).Scan(&response)
	return response, result.Error
}

func (repo VendorRepository) VendorExists(name string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from vendors where name = ?)", name).Scan(&exists)
	return exists
}

func (repo VendorRepository) VendorExistsByID(id string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from vendors where id = ?)", id).Scan(&exists)
	return exists
}

func (repo VendorRepository) Update(id string, Vendor *entities.Vendor) error {
	result := repo.db.Table("vendors").Where("id = ?", id).Updates(Vendor)
	return result.Error
}

func (repo VendorRepository) Delete(id string) error {
	result := repo.db.Table("vendors").Unscoped().Where("id = ?", id).Delete(&entities.Vendor{})
	return result.Error
}
