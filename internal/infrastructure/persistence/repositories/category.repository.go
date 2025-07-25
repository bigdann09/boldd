package repositories

import (
	"fmt"
	"strings"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/pkgs/utils"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Delete(id string) error
	CategoryExists(name string) bool
	CategoryExistsByID(id string) bool
	Create(address *entities.Category) error
	Find(id string) (*dtos.CategoryResponse, error)
	Update(id string, category *entities.Category) error
	FindAllPaginated(filter *dtos.CategoryQueryFilter) (utils.PaginationResponse[dtos.CategoryResponse], error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (repo CategoryRepository) Create(address *entities.Category) error {
	result := repo.db.Table("categories").Create(&address)
	return result.Error
}

func (repo CategoryRepository) FindAllPaginated(filter *dtos.CategoryQueryFilter) (utils.PaginationResponse[dtos.CategoryResponse], error) {
	if strings.EqualFold(filter.SortBy, "") {
		filter.SortBy = "name"
	}

	if strings.EqualFold(filter.Order, "") {
		filter.Order = "asc"
	}

	query := repo.db.Table("categories").Order(fmt.Sprintf("%s %s", filter.SortBy, filter.Order))
	return utils.NewPaginationResponse[dtos.CategoryResponse](filter.Page, filter.PageSize, query)
}

func (repo CategoryRepository) Find(id string) (*dtos.CategoryResponse, error) {
	var response *dtos.CategoryResponse
	result := repo.db.Table("categories").Where("id = ?", id).Scan(&response)
	return response, result.Error
}

func (repo CategoryRepository) CategoryExists(name string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from categories where name = ?)", name).Scan(&exists)
	return exists
}

func (repo CategoryRepository) CategoryExistsByID(id string) bool {
	var exists bool
	repo.db.Raw("select exists (select 1 from categories where id = ?)", id).Scan(&exists)
	return exists
}

func (repo CategoryRepository) Update(id string, category *entities.Category) error {
	result := repo.db.Table("categories").Where("id = ?", id).Updates(category)
	return result.Error
}

func (repo CategoryRepository) Delete(id string) error {
	result := repo.db.Table("categories").Unscoped().Where("id = ?", id).Delete(&entities.Category{})
	return result.Error
}
