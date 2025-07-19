package repositories

import (
	"fmt"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/pkgs/utils"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Create(address *entities.Category) error
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
	query := repo.db.Table("categories").Order(fmt.Sprintf("%s %s", filter.SortBy, filter.Order))
	return utils.NewPaginationResponse[dtos.CategoryResponse](filter.Page, filter.PageSize, query)
}
