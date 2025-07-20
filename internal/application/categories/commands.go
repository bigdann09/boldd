package categories

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"go.uber.org/zap"
)

type ICategoryCommand interface {
	Delete(uuid string) interface{}
	Create(payload *CreateCategoryRequest) interface{}
	Update(uuid string, payload *UpdateCategoryRequest) interface{}
}

type CategoryCommand struct {
	logger             *zap.Logger
	categoryRepository repositories.ICategoryRepository
	categoryCache      *cache.Cache[entities.Category]
}

func NewCategoryCommand(
	logger *zap.Logger,
	categoryRepository repositories.ICategoryRepository,
	categoryCache *cache.Cache[entities.Category],
) *CategoryCommand {
	return &CategoryCommand{logger, categoryRepository, categoryCache}
}

func (cmd CategoryCommand) Create(payload *CreateCategoryRequest) interface{} {
	cmd.logger.Info("check if category already exist")
	if exists := cmd.categoryRepository.CategoryExists(payload.Name); exists {
		cmd.logger.Warn("category already exists/stored", zap.String("category", payload.Name))
		return dtos.ErrorResponse{Message: "category already stored", Status: http.StatusBadRequest}
	}

	cmd.logger.Info("storing a new category to store")
	err := cmd.categoryRepository.Create(entities.NewCategory(payload.Name))
	if err != nil {
		fmt.Println(err)
		cmd.logger.Error("encountered an error stroing category", zap.Error(err))
		return dtos.ErrorResponse{Message: "could not create category", Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidating cache")
	cmd.categoryCache.Delete("categories:all")

	return nil
}

func (cmd CategoryCommand) Delete(uuid string) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("uuid", uuid))
	if exists := cmd.categoryRepository.CategoryExistsByUUID(uuid); !exists {
		cmd.logger.Warn("category record not found", zap.String("uuid", uuid))
		return dtos.ErrorResponse{Message: "category not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("deleting category record", zap.String("uuid", uuid))
	err := cmd.categoryRepository.Delete(uuid)
	if err != nil {
		cmd.logger.Error("could not delete category record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.categoryCache.Delete("categories:all")
	return nil
}

func (cmd CategoryCommand) Update(uuid string, payload *UpdateCategoryRequest) interface{} {
	cmd.logger.Info("check if record exists before updating", zap.String("uuid", uuid))
	if exists := cmd.categoryRepository.CategoryExistsByUUID(uuid); !exists {
		cmd.logger.Warn("category record not found", zap.String("uuid", uuid))
		return dtos.ErrorResponse{Message: "category not found", Status: http.StatusNotFound}
	}

	cmd.logger.Info("updating category record", zap.String("uuid", uuid))
	err := cmd.categoryRepository.Update(uuid, &entities.Category{Name: payload.Name})
	if err != nil {
		cmd.logger.Error("could not update category record", zap.Error(err))
		return dtos.ErrorResponse{Message: err.Error(), Status: http.StatusInternalServerError}
	}

	cmd.logger.Info("invalidate cache")
	cmd.categoryCache.Delete("categories:all")
	return nil
}
