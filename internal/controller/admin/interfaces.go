package admin

import (
	"context"
	"referalMS/internal/controller/dto"
	"referalMS/internal/domain/entity"
)

type AdminService interface {
	GetWinners(ctx context.Context, dto dto.ExternalAdminDTO, filters dto.GetWinnersDTO) (winners []entity.Referal, err error)
	RegisterNewAdmin(ctx context.Context, adminDTO dto.AdminDTO) (adminId int64, err error)
}
