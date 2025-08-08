package attributes

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type IAttributeQuery interface {
	Get(id string) (*dtos.AttributeResponse, interface{})
	GetAll(filter *dtos.AttributeQueryFilter) (utils.PaginationResponse[dtos.AttributeResponse], interface{})
}

type AttributeQuery struct {
	logger              *zap.Logger
	attributeRepository repositories.IAttributeRepository
	attributeCache      *cache.Cache[*dtos.AttributeResponse]
	attributeAllCache   *cache.Cache[utils.PaginationResponse[dtos.AttributeResponse]]
}

func NewAttributeQuery(
	logger *zap.Logger,
	attributeRepository repositories.IAttributeRepository,
	attributeCache *cache.Cache[*dtos.AttributeResponse],
	attributeAllCache *cache.Cache[utils.PaginationResponse[dtos.AttributeResponse]],
) *AttributeQuery {
	return &AttributeQuery{logger, attributeRepository, attributeCache, attributeAllCache}
}

func (qry *AttributeQuery) GetAll(filter *dtos.AttributeQueryFilter) (utils.PaginationResponse[dtos.AttributeResponse], interface{}) {
	key := "attributes:all"
	if filter.Page > 0 {
		key = fmt.Sprintf(
			"attributes_page%d_size%d_sortby%s_order%s",
			filter.Page, filter.PageSize, filter.SortBy, filter.Order,
		)
	}

	fmt.Println("Attribute key", key)

	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	attributes, err := qry.attributeAllCache.GetOrSet(
		key,
		func() (utils.PaginationResponse[dtos.AttributeResponse], error) {
			return qry.attributeRepository.FindAllPaginated(filter)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving attributes from cache", zap.Error(err))
		return utils.PaginationResponse[dtos.AttributeResponse]{}, dtos.ErrorResponse{Message: "could not fetch attributes", Status: http.StatusInternalServerError}
	}
	return attributes, nil
}

func (qry *AttributeQuery) Get(id string) (*dtos.AttributeResponse, interface{}) {
	key := fmt.Sprintf("attribute_%s", id)
	qry.logger.Info("retrieving from cache if data exists else setting to cache")
	attribute, err := qry.attributeCache.GetOrSet(
		key,
		func() (*dtos.AttributeResponse, error) {
			return qry.attributeRepository.Find(id)
		},
	)
	if err != nil {
		qry.logger.Error("error retrieving attribute from cache", zap.Error(err))
		return &dtos.AttributeResponse{}, dtos.ErrorResponse{Message: "could not fetch attributes", Status: http.StatusInternalServerError}
	}

	return attribute, nil
}
