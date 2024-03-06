package admin

import (
	"context"
	"referalMS/internal/domain/entity"
)

type ReadAdminStorage interface {
	GetAdmin(ctx context.Context, login, token string) (admin entity.Admin, err error)
	GetWinners(ctx context.Context, admin entity.Admin, winnersFilter entity.WinnersFilter) (winners []entity.Referal, err error)
}

type CreateAdminUseCase interface {
	Execute(ctx context.Context, admin entity.Admin) (adminId int64, err error)
}
