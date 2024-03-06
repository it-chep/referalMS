package create_admin

import (
	"context"
	"log/slog"
	"referalMS/internal/domain/entity"
)

type CreateAdminUseCase struct {
	repo   WriteRepo
	logger *slog.Logger
}

func NewCreateAdminUseCase(repo WriteRepo, logger *slog.Logger) *CreateAdminUseCase {
	return &CreateAdminUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *CreateAdminUseCase) Execute(ctx context.Context, admin entity.Admin) (adminId int64, err error) {
	u.logger.Info("23232323")
	createAdmin, err := u.repo.CreateAdmin(ctx, admin)
	u.logger.Info("CreateAdminUseCase) Execute", createAdmin, err)
	if err != nil {
		return 0, err
	}
	return createAdmin, nil
}
