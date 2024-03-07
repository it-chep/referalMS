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
	return u.repo.CreateAdmin(ctx, admin)
}
