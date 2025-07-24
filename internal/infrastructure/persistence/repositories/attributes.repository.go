package repositories

import (
	"fmt"
	"strings"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/pkgs/utils"
	"gorm.io/gorm"
)

type IAttributeRepository interface {
	Delete(id string) error
	AttributeExists(name string) bool
	AttributeExistsByID(id string) bool
	Create(address *entities.Attribute) error
	Find(id string) (*dtos.AttributeResponse, error)
	Update(id string, Attribute *entities.Attribute) error
	FindAllPaginated(filter *dtos.AttributeQueryFilter) (utils.PaginationResponse[dtos.AttributeResponse], error)
}

type AttributeRepository struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) *AttributeRepository {
	return &AttributeRepository{db}
}

func (repo AttributeRepository) Create(address *entities.Attribute) error {
	result := repo.db.Table("attributes").Create(&address)
	return result.Error
}

func (repo AttributeRepository) FindAllPaginated(filter *dtos.AttributeQueryFilter) (utils.PaginationResponse[dtos.AttributeResponse], error) {
	if strings.EqualFold(filter.SortBy, "") {
		filter.SortBy = "name"
	}

	if strings.EqualFold(filter.Order, "") {
		filter.Order = "asc"
	}

	query := repo.db.Table("attributes").Order(fmt.Sprintf("%s %s", filter.SortBy, filter.Order))
	return utils.NewPaginationResponse[dtos.AttributeResponse](filter.Page, filter.PageSize, query)
}

func (repo AttributeRepository) Find(id string) (*dtos.AttributeResponse, error) {
	var response *dtos.AttributeResponse
	result := repo.db.Table("attributes").Where("id = ?", id).Scan(&response)
	return response, result.Error
}

func (repo AttributeRepository) AttributeExists(name string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from attributes where name = ?)", name).Scan(&exists)
	return exists
}

func (repo AttributeRepository) AttributeExistsByID(id string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from attributes where id = ?)", id).Scan(&exists)
	return exists
}

func (repo AttributeRepository) Update(id string, Attribute *entities.Attribute) error {
	result := repo.db.Table("attributes").Where("id = ?", id).Updates(Attribute)
	return result.Error
}

func (repo AttributeRepository) Delete(id string) error {
	result := repo.db.Table("attributes").Unscoped().Where("id = ?", id).Delete(&entities.Attribute{})
	return result.Error
}
