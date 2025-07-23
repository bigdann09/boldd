package profile

import (
	"net/http"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type IProfileCommand interface {
	ChangePassword(user *dtos.UserResponse, payload *ChangePasswordRequest) interface{}
}

type ProfileCommand struct {
	logger         *zap.Logger
	userRepository repositories.IUserRepository
	cacheUser      *cache.Cache[*dtos.UserResponse]
}

func NewProfileCommand(logger *zap.Logger, userRepository repositories.IUserRepository, cacheUser *cache.Cache[*dtos.UserResponse]) *ProfileCommand {
	return &ProfileCommand{logger, userRepository, cacheUser}
}

func (cmd ProfileCommand) ChangePassword(user *dtos.UserResponse, payload *ChangePasswordRequest) interface{} {
	cmd.logger.Info("compare old password to see if it matches")
	if err := utils.ComparePasswords(user.Password, payload.OldPassword); err != nil {
		cmd.logger.Error("old password do not match with current password", zap.Error(err))
		return dtos.ErrorResponse{Status: http.StatusBadRequest, Message: "provided password is incorrect"}
	}

	cmd.logger.Info("update user password")
	if err := cmd.userRepository.Update(user.ID, &entities.User{Password: utils.HashPassword(payload.NewPassword)}); err != nil {
		cmd.logger.Error("there was an error updating user password", zap.Error(err))
		return dtos.ErrorResponse{Status: http.StatusInternalServerError, Message: "there was an error updating password"}
	}

	return nil
}
