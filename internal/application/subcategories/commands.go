package subcategories

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"go.uber.org/zap"
)

type ISubCategoryCommand interface {
	Delete(uuid string) interface{}
	Create(payload *CreateSubCategoryRequest) interface{}
	Update(uuid string, payload *UpdateSubCategoryRequest) interface{}
}

type SubCategoryCommand struct {
	logger                *zap.Logger
	categoryRepository    repositories.ICategoryRepository
	subcategoryRepository repositories.ISubCategoryRepository
	subcategoryCache      *cache.Cache[entities.SubCategory]
}

func NewSubCategoryCommand(
	logger *zap.Logger,
	categoryRepository repositories.ICategoryRepository,
	subcategoryRepository repositories.ISubCategoryRepository,
	subcategoryCache *cache.Cache[entities.SubCategory],
) *SubCategoryCommand {
	return &SubCategoryCommand{logger, categoryRepository, subcategoryRepository, subcategoryCache}
}

func (cmd SubCategoryCommand) Create(payload *CreateSubCategoryRequest) interface{} {
	cmd.logger.Info("check if category ID is registered or valid")
	if exists := cmd.categoryRepository.CategoryExistsByID(payload.CategoryID); !exists {
		cmd.logger.Warn("category record not found or invalid", zap.Uint("category", payload.CategoryID))
		return dtos.ErrorResponse{Message: "category record not found or invalid", Status: http.StatusBadRequest}
	}

	cmd.logger.Info("check if subcategory already exist")
	if exists := cmd.subcategoryRepository.SubCategoryExists(payload.Name); exists {
		cmd.logger.Warn("subcategory already exists/stored", zap.String("subcategory", payload.Name))
		return dtos.ErrorResponse{Message: "subcategory already stored", Status: http.StatusBadRequest}
	}

	cmd.logger.Info("storing a new category to store")
	err := cmd.subcategoryRepository.Create(entities.NewSubCategory(payload.Name, payload.CategoryID))
	if err != nil {
		fmt.Println(err)
		cmd.logger.Error("encountered an error stroing subcategory", zap.Error(err))
		return dtos.ErrorResponse{Message: "could not create subcategory", Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidating cache")
	cmd.subcategoryCache.Delete("subcategories:all")

	return nil
}

func (cmd SubCategoryCommand) Delete(uuid string) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("uuid", uuid))
	if exists := cmd.subcategoryRepository.SubCategoryExistsByUUID(uuid); !exists {
		cmd.logger.Warn("subcategory record not found", zap.String("uuid", uuid))
		return dtos.ErrorResponse{Message: "subcategory not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("deleting subcategory record", zap.String("uuid", uuid))
	err := cmd.subcategoryRepository.Delete(uuid)
	if err != nil {
		cmd.logger.Error("could not delete subcategory record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.subcategoryCache.Delete("subcategories:all")
	return nil
}

func (cmd SubCategoryCommand) Update(uuid string, payload *UpdateSubCategoryRequest) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("uuid", uuid))
	if exists := cmd.subcategoryRepository.SubCategoryExistsByUUID(uuid); !exists {
		cmd.logger.Warn("subcategory record not found", zap.String("uuid", uuid))
		return dtos.ErrorResponse{Message: "subcategory not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("updating subcategory record", zap.String("uuid", uuid))
	err := cmd.subcategoryRepository.Update(uuid, entities.UpdateSubCategory(payload.Name))
	if err != nil {
		cmd.logger.Error("could not update subcategory record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.subcategoryCache.Delete("subcategories:all")
	return nil
}
