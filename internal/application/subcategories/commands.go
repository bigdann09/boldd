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
	Delete(id string) interface{}
	Create(payload *CreateSubCategoryRequest) interface{}
	Update(id string, payload *UpdateSubCategoryRequest) interface{}
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
		cmd.logger.Warn("category record not found or invalid", zap.String("category", payload.CategoryID))
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

func (cmd SubCategoryCommand) Delete(id string) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("id", id))
	if exists := cmd.subcategoryRepository.SubCategoryExistsByID(id); !exists {
		cmd.logger.Warn("subcategory record not found", zap.String("id", id))
		return dtos.ErrorResponse{Message: "subcategory not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("deleting subcategory record", zap.String("id", id))
	err := cmd.subcategoryRepository.Delete(id)
	if err != nil {
		cmd.logger.Error("could not delete subcategory record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.subcategoryCache.Delete("subcategories:all")
	return nil
}

func (cmd SubCategoryCommand) Update(id string, payload *UpdateSubCategoryRequest) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("id", id))
	if exists := cmd.subcategoryRepository.SubCategoryExistsByID(id); !exists {
		cmd.logger.Warn("subcategory record not found", zap.String("id", id))
		return dtos.ErrorResponse{Message: "subcategory not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("updating subcategory record", zap.String("id", id))
	err := cmd.subcategoryRepository.Update(id, entities.UpdateSubCategory(payload.Name))
	if err != nil {
		cmd.logger.Error("could not update subcategory record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.subcategoryCache.Delete("subcategories:all")
	return nil
}
