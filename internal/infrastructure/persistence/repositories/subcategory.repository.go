package repositories

import (
	"fmt"
	"strings"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/pkgs/utils"
	"gorm.io/gorm"
)

type ISubCategoryRepository interface {
	Delete(uuid string) error
	SubCategoryExists(name string) bool
	SubCategoryExistsByID(uuid string) bool
	Create(address *entities.SubCategory) error
	Find(uuid string) (*dtos.SubCategoryResponse, error)
	Update(uuid string, category *entities.SubCategory) error
	FindAllPaginated(filter *dtos.SubCategoryQueryFilter) (utils.PaginationResponse[dtos.SubCategoryResponse], error)
}

type SubCategoryRepository struct {
	db *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) *SubCategoryRepository {
	return &SubCategoryRepository{db}
}

func (repo SubCategoryRepository) Create(address *entities.SubCategory) error {
	result := repo.db.Table("subcategories").Create(&address)
	return result.Error
}

func (repo SubCategoryRepository) FindAllPaginated(filter *dtos.SubCategoryQueryFilter) (utils.PaginationResponse[dtos.SubCategoryResponse], error) {
	if strings.EqualFold(filter.SortBy, "") {
		filter.SortBy = "name"
	}

	if strings.EqualFold(filter.Order, "") {
		filter.Order = "asc"
	}

	query := repo.db.Table("subcategories").Order(fmt.Sprintf("%s %s", filter.SortBy, filter.Order))
	return utils.NewPaginationResponse[dtos.SubCategoryResponse](filter.Page, filter.PageSize, query)
}

func (repo SubCategoryRepository) Find(id string) (*dtos.SubCategoryResponse, error) {
	var response *dtos.SubCategoryResponse
	result := repo.db.Table("subcategories").Where("id = ?", id).Scan(&response)
	return response, result.Error
}

func (repo SubCategoryRepository) SubCategoryExists(name string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from subcategories where name = ?)", name).Scan(&exists)
	return exists
}

func (repo SubCategoryRepository) SubCategoryExistsByID(id string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from subcategories where id = ?)", id).Scan(&exists)
	return exists
}

func (repo SubCategoryRepository) Update(id string, category *entities.SubCategory) error {
	result := repo.db.Table("subcategories").Where("id = ?", id).Updates(category)
	return result.Error
}

func (repo SubCategoryRepository) Delete(id string) error {
	result := repo.db.Table("subcategories").Unscoped().Where("id = ?", id).Delete(&entities.SubCategory{})
	return result.Error
}
