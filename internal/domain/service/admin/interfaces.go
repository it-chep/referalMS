package admin

import (
	"context"
	"referalMS/internal/domain/entity"
)

type ReadAdminStorage interface {
	GetAdmin(ctx context.Context, login, password, token string) (adminID int64, err error)
	GetWinners(ctx context.Context, admin entity.Admin, winnersFilter entity.WinnersFilter) (winners []entity.Referal, err error)
}
