package referal

import (
	"context"
	"referalMS/internal/domain/entity"
)

type ReadReferalStorage interface {
	GetReferalByID(ctx context.Context, ID, inServiceId int64, admin entity.Admin) (referal entity.Referal, err error)
	GetReferalByTgID(ctx context.Context, tgID, inServiceId int64, admin entity.Admin) (referal entity.Referal, err error)
	GetReferalByRefLink(ctx context.Context, referalLink string, admin entity.Admin) (referal entity.Referal, err error)
	GetReferalStatistic(ctx context.Context, tgID int64, admin entity.Admin, daysAgo int) (allUsers int64, lastNDays int64, err error)
}

type CreateReferalUseCase interface {
	Execute(ctx context.Context, referal entity.Referal, admin entity.Admin) (referalId int64, err error)
}

//type ReadReferalAdminStorage interface {
//	GetAdmin(ctx context.Context, dto dto.ExternalAdminDTO) (admin entity.Admin)
//}
