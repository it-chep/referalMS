package user

import (
	"context"
	"fmt"
	"log/slog"
	"referalMS/internal/controller/dto"
	"referalMS/internal/domain/entity"
	"referalMS/internal/domain/service/admin"
)

type UserService struct {
	createUserUseCase CreateUserUseCase
	adminService      admin.AdminService
	readRepo          ReadUserStorage
	logger            *slog.Logger
}

func NewUserService(
	createUserUseCase CreateUserUseCase,
	readRepo ReadUserStorage,
	adminService admin.AdminService,
	logger *slog.Logger,
) UserService {
	return UserService{
		createUserUseCase: createUserUseCase,
		readRepo:          readRepo,
		adminService:      adminService,
		logger:            logger,
	}
}

func (u UserService) RegisterNewUser(ctx context.Context, dto dto.TgUserDTO, adto dto.ExternalAdminDTO) (userId int64, err error) {
	adminEntity, err := u.adminService.GetAdmin(ctx, adto)
	if err != nil {
		return 0, fmt.Errorf("adminEntity not found")
	}
	u.logger.Info("adminEntity was found, trying to register new referal")

	user, err := u.readRepo.GetUserByReferalLink(ctx, dto.ReferalLink)
	if err != nil {
		return 0, err
	}

	if user != (entity.User{}) {
		return 0, fmt.Errorf("user has registered")
	}

	newUser := dto.ToDomain(adminEntity.GetId())

	// Save new referral
	userId, err = u.createUserUseCase.Execute(ctx, newUser, adminEntity)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
