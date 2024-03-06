package create_referal

import (
	"context"
	"referalMS/internal/domain/entity"
)

type CreateReferalUseCase struct {
	repo WriteRepo
}

func NewCreateReferalUseCase(repo WriteRepo) *CreateReferalUseCase {
	return &CreateReferalUseCase{
		repo: repo,
	}
}

func (uc CreateReferalUseCase) Execute(ctx context.Context, referal entity.Referal, admin entity.Admin) (referalId int64, err error) {
	return uc.repo.CreateReferal(ctx, referal, admin)
}
