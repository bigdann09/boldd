package categories

import (
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
)

type ICategoryQuery interface {
}

type CategoryQuery struct {
	categoryRepository repositories.ICategoryRepository
}

func NewCategoryQuery(categoryRepository repositories.ICategoryRepository) *CategoryQuery {
	return &CategoryQuery{categoryRepository}
}

func (qry *CategoryQuery) GetAll(filter *dtos.CategoryQueryFilter) (utils.PaginationResponse[dtos.CategoryResponse], interface{}) {
	categories, err := qry.categoryRepository.FindAllPaginated(filter)
	if err != nil {
		return utils.PaginationResponse[dtos.CategoryResponse]{}, map[string]interface{}{"error": "could not fetch categories", "code": http.StatusInternalServerError}
	}
	return categories, nil
}
