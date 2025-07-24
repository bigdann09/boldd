package attributes

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"go.uber.org/zap"
)

type IAttributeCommand interface {
	Delete(id string) interface{}
	Create(payload *CreateAttributeRequest) interface{}
	Update(id string, payload *UpdateAttributeRequest) interface{}
}

type AttributeCommand struct {
	logger              *zap.Logger
	attributeRepository repositories.IAttributeRepository
	attributeCache      *cache.Cache[entities.Attribute]
}

func NewAttributeCommand(
	logger *zap.Logger,
	attributeRepository repositories.IAttributeRepository,
	attributeCache *cache.Cache[entities.Attribute],
) *AttributeCommand {
	return &AttributeCommand{logger, attributeRepository, attributeCache}
}

func (cmd AttributeCommand) Create(payload *CreateAttributeRequest) interface{} {
	cmd.logger.Info("check if category already exist")
	if exists := cmd.attributeRepository.AttributeExists(payload.Name); exists {
		cmd.logger.Warn("attribute already exists/stored", zap.String("attribute", payload.Name))
		return dtos.ErrorResponse{Message: "attribute already stored", Status: http.StatusBadRequest}
	}

	cmd.logger.Info("storing a new attribute to store")
	err := cmd.attributeRepository.Create(entities.NewAttribute(payload.Name))
	if err != nil {
		fmt.Println(err)
		cmd.logger.Error("encountered an error storing attribute", zap.Error(err))
		return dtos.ErrorResponse{Message: "could not create attribute", Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidating cache")
	cmd.attributeCache.Delete("attributes:all")
	return nil
}

func (cmd AttributeCommand) Delete(id string) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("id", id))
	if exists := cmd.attributeRepository.AttributeExistsByID(id); !exists {
		cmd.logger.Warn("attribute record not found", zap.String("id", id))
		return dtos.ErrorResponse{Message: "attribute not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("deleting attribute record", zap.String("id", id))
	err := cmd.attributeRepository.Delete(id)
	if err != nil {
		cmd.logger.Error("could not delete attribute record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.attributeCache.Delete("attributes:all")
	return nil
}

func (cmd AttributeCommand) Update(id string, payload *UpdateAttributeRequest) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("id", id))
	if exists := cmd.attributeRepository.AttributeExistsByID(id); !exists {
		cmd.logger.Warn("attribute record not found", zap.String("id", id))
		return dtos.ErrorResponse{Message: "attribute not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("updating attribute record", zap.String("id", id))
	err := cmd.attributeRepository.Update(id, entities.UpdateAttribute(payload.Name))
	if err != nil {
		cmd.logger.Error("could not update attribute record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.attributeCache.Delete("attributes:all")
	return nil
}
