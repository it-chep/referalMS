package create_referal

import (
	"context"
	"log/slog"
	"referalMS/internal/domain/entity"
)

type CreateReferalUseCase struct {
	repo   WriteRepo
	logger *slog.Logger
}

func NewCreateReferalUseCase(repo WriteRepo, logger *slog.Logger) *CreateReferalUseCase {
	return &CreateReferalUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc CreateReferalUseCase) Execute(ctx context.Context, referal entity.Referal, admin entity.Admin) (referalId int64, err error) {
	return uc.repo.CreateReferal(ctx, referal, admin)
}
