package categories

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type ICategoryQuery interface {
	Get(id string) (*dtos.CategoryResponse, interface{})
	GetAll(filter *dtos.CategoryQueryFilter) (utils.PaginationResponse[dtos.CategoryResponse], interface{})
}

type CategoryQuery struct {
	logger             *zap.Logger
	categoryRepository repositories.ICategoryRepository
	categoryCache      *cache.Cache[*dtos.CategoryResponse]
	categoryAllCache   *cache.Cache[utils.PaginationResponse[dtos.CategoryResponse]]
}

func NewCategoryQuery(
	logger *zap.Logger,
	categoryRepository repositories.ICategoryRepository,
	categoryCache *cache.Cache[*dtos.CategoryResponse],
	categoryAllCache *cache.Cache[utils.PaginationResponse[dtos.CategoryResponse]],
) *CategoryQuery {
	return &CategoryQuery{logger, categoryRepository, categoryCache, categoryAllCache}
}

func (qry *CategoryQuery) GetAll(filter *dtos.CategoryQueryFilter) (utils.PaginationResponse[dtos.CategoryResponse], interface{}) {
	key := "categories:all"
	if filter.Page > 0 {
		key = fmt.Sprintf(
			"categories_page%d_size%d_sortby%s_order%s",
			filter.Page, filter.PageSize, filter.SortBy, filter.Order,
		)
	}

	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	categories, err := qry.categoryAllCache.GetOrSet(
		key,
		func() (utils.PaginationResponse[dtos.CategoryResponse], error) {
			return qry.categoryRepository.FindAllPaginated(filter)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving categories from cache", zap.Error(err))
		return utils.PaginationResponse[dtos.CategoryResponse]{}, dtos.ErrorResponse{Message: "could not fetch categories", Status: http.StatusInternalServerError}
	}
	return categories, nil
}

func (qry *CategoryQuery) Get(id string) (*dtos.CategoryResponse, interface{}) {
	key := fmt.Sprintf("categories_%s", id)
	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	category, err := qry.categoryCache.GetOrSet(
		key,
		func() (*dtos.CategoryResponse, error) {
			return qry.categoryRepository.Find(id)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving categories from cache", zap.Error(err))
		return &dtos.CategoryResponse{}, dtos.ErrorResponse{Message: "could not fetch categories", Status: http.StatusInternalServerError}
	}

	return category, nil
}
