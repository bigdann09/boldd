package vendors

import (
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"go.uber.org/zap"
)

type IVendorCommand interface {
	// Delete(id string) interface{}
	Create(user *dtos.UserResponse, payload *CreateVendorRequest) interface{}
	// Update(id string, payload *UpdateVendorRequest) interface{}
}

type VendorCommand struct {
	logger           *zap.Logger
	vendorRepository repositories.IVendorRepository
	vendorCache      *cache.Cache[dtos.VendorResponse]
}

func NewVendorCommand(
	logger *zap.Logger,
	vendorRepository repositories.IVendorRepository,
	vendorCache *cache.Cache[dtos.VendorResponse],
) *VendorCommand {
	return &VendorCommand{logger: logger, vendorRepository: vendorRepository, vendorCache: vendorCache}
}

func (srv VendorCommand) Create(user *dtos.UserResponse, payload *CreateVendorRequest) interface{} {
	srv.logger.Info("check if vendor already exists")
	if exists := srv.vendorRepository.VendorExists(payload.Name); exists {
		srv.logger.Warn("vendor already exists/stored", zap.String("vendor", payload.Name))
		return dtos.ErrorResponse{Message: "vendor already stored", Status: http.StatusBadRequest}
	}

	srv.logger.Info("storing a new vendor to store")
	err := srv.vendorRepository.Create(entities.NewVendor(
		int(user.ID),
		payload.Name,
		payload.BusinessEmail,
		payload.BusinessAddress,
		payload.BusinessPhone,
		payload.Description,
	))
	if err != nil {
		srv.logger.Error("encountered an error storing vendor", zap.Error(err))
		return dtos.ErrorResponse{Message: "could not create vendor", Status: http.StatusInternalServerError}
	}

	srv.logger.Info("invalidating cache")
	srv.vendorCache.Delete("vendors:all")

	return nil
}
