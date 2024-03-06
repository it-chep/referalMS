package v1

import (
	"context"
	"referalMS/internal/controller/dto"
)

type UserService interface {
	RegisterNewUser(ctx context.Context, dto dto.TgUserDTO, adto dto.ExternalAdminDTO) (userId int64, err error)
}

type ReferalService interface {
	RegisterNewReferal(ctx context.Context, dto dto.ReferalUserDTO, adto dto.ExternalAdminDTO) (referalLink string, err error)
	GetReferalStatistic(ctx context.Context, dto dto.ReferalStatisticDTO, adto dto.ExternalAdminDTO) (allUsers, lastNDays int64, err error)
}
