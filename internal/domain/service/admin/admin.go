package admin

import (
	"context"
	"fmt"
	"referalMS/internal/controller/dto"
	"referalMS/internal/domain/entity"
)

type AdminService struct {
	ReadRepo ReadAdminStorage
}

func NewAdminService(readRepo ReadAdminStorage) AdminService {
	return AdminService{
		ReadRepo: readRepo,
	}
}

func (s *AdminService) GetAdmin(ctx context.Context, dto dto.ExternalAdminDTO) (admin entity.Admin) {
	adminId, err := s.ReadRepo.GetAdmin(
		ctx,
		dto.Login,
		dto.Password,
		dto.Token,
	)
	if err != nil {
		return entity.Admin{}
	}
	admin = *entity.NewAdmin(
		dto.Login, dto.Password, dto.Token, entity.WithAdmId(adminId),
	)
	return admin
}

func (s *AdminService) GetWinners(ctx context.Context, dto dto.ExternalAdminDTO, filters dto.GetWinnersDTO) (winners []entity.Referal, err error) {
	admin := s.GetAdmin(ctx, dto)
	if admin == (entity.Admin{}) {
		return nil, fmt.Errorf("admin not found")
	}

	winnersFilter := entity.NewWinnersFilter(
		filters.Limit,
		filters.DaysInterval,
	)

	winners, err = s.ReadRepo.GetWinners(ctx, admin, *winnersFilter)
	if err != nil {
		return nil, err
	}

	return winners, nil
}
