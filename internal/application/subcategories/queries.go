package subcategories

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type ISubCategoryQuery interface {
	Get(id string) (*dtos.SubCategoryResponse, interface{})
	GetAll(filter *dtos.SubCategoryQueryFilter) (utils.PaginationResponse[dtos.SubCategoryResponse], interface{})
}

type SubCategoryQuery struct {
	logger                *zap.Logger
	subcategoryRepository repositories.ISubCategoryRepository
	subcategoryCache      *cache.Cache[*dtos.SubCategoryResponse]
	subcategoryAllCache   *cache.Cache[utils.PaginationResponse[dtos.SubCategoryResponse]]
}

func NewSubCategoryQuery(
	logger *zap.Logger,
	subcategoryRepository repositories.ISubCategoryRepository,
	subcategoryCache *cache.Cache[*dtos.SubCategoryResponse],
	subcategoryAllCache *cache.Cache[utils.PaginationResponse[dtos.SubCategoryResponse]],
) *SubCategoryQuery {
	return &SubCategoryQuery{logger, subcategoryRepository, subcategoryCache, subcategoryAllCache}
}

func (qry *SubCategoryQuery) GetAll(filter *dtos.SubCategoryQueryFilter) (utils.PaginationResponse[dtos.SubCategoryResponse], interface{}) {
	key := "subcategories:all"
	if filter.Page > 0 {
		key = fmt.Sprintf(
			"subcategories_page%d_size%d_sortby%s_order%s",
			filter.Page, filter.PageSize, filter.SortBy, filter.Order,
		)
	}

	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	subcategories, err := qry.subcategoryAllCache.GetOrSet(
		key,
		func() (utils.PaginationResponse[dtos.SubCategoryResponse], error) {
			return qry.subcategoryRepository.FindAllPaginated(filter)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving subcategories from cache", zap.Error(err))
		return utils.PaginationResponse[dtos.SubCategoryResponse]{}, dtos.ErrorResponse{Message: "could not fetch subcategories", Status: http.StatusInternalServerError}
	}
	return subcategories, nil
}

func (qry *SubCategoryQuery) Get(id string) (*dtos.SubCategoryResponse, interface{}) {
	key := fmt.Sprintf("subcategories_%s", id)
	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	subcategory, err := qry.subcategoryCache.GetOrSet(
		key,
		func() (*dtos.SubCategoryResponse, error) {
			return qry.subcategoryRepository.Find(id)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving subcategories from cache", zap.Error(err))
		return &dtos.SubCategoryResponse{}, dtos.ErrorResponse{Message: "could not fetch subcategories", Status: http.StatusInternalServerError}
	}

	return subcategory, nil
}
