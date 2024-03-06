package referal

import (
	"context"
	"fmt"
	"log/slog"
	"referalMS/internal/controller/dto"
	"referalMS/internal/domain/entity"
	"referalMS/internal/domain/service/admin"
	"strconv"
	"time"
)

type ReferalService struct {
	CreateReferalUseCase CreateReferalUseCase
	ReadRepo             ReadReferalStorage
	AdminService         admin.AdminService
	logger               *slog.Logger
}

func NewReferalService(
	createReferalUseCase CreateReferalUseCase,
	readRepo ReadReferalStorage,
	adminService admin.AdminService,
	logger *slog.Logger,
) ReferalService {

	return ReferalService{
		CreateReferalUseCase: createReferalUseCase,
		ReadRepo:             readRepo,
		AdminService:         adminService,
		logger:               logger,
	}
}

func (s *ReferalService) RegisterNewReferal(ctx context.Context, dto dto.ReferalUserDTO, adto dto.ExternalAdminDTO) (referalLink string, err error) {
	adminEntity, err := s.AdminService.GetAdmin(ctx, adto)
	if err != nil {
		return "", fmt.Errorf("adminEntity not found")
	}
	s.logger.Info("adminEntity was found, trying to register new referal")

	referal, err := s.ReadRepo.GetReferalByTgID(ctx, dto.TgId, dto.InServiceID, adminEntity)
	if err != nil {
		return "", err
	}

	if referal != (entity.Referal{}) {
		return "", fmt.Errorf("referal has registered")
	}

	referalLink = s.generateReferalLink(
		dto.TgId,
		dto.InServiceID,
		adminEntity.GetId(),
	)
	s.logger.Info("link was generated successfully")

	newReferral := dto.ToDomain(adminEntity.GetId())

	// Save new referral
	_, err = s.CreateReferalUseCase.Execute(ctx, *newReferral, adminEntity)
	if err != nil {
		return "", err
	}

	return referalLink, nil
}

func (s *ReferalService) GetReferalStatistic(ctx context.Context, dto dto.ReferalStatisticDTO, adto dto.ExternalAdminDTO) (allUsers, lastNDays int64, err error) {
	adminEntity, err := s.AdminService.GetAdmin(ctx, adto)
	if err != nil {
		return 0, 0, fmt.Errorf("adminEntity not found")
	}
	s.logger.Info("adminEntity was found, trying to collect referal statistic")
	// Todo add default daysAgo from config
	allUsers, lastNDays, err = s.ReadRepo.GetReferalStatistic(ctx, dto.TgId, adminEntity, 30)
	if err != nil {
		return 0, 0, err
	}

	return allUsers, lastNDays, nil
}

func (s *ReferalService) GetReferalByRefLink(ctx context.Context, referalLink string, adto dto.ExternalAdminDTO) (referal entity.Referal, err error) {
	adminEntity, err := s.AdminService.GetAdmin(ctx, adto)
	if err != nil {
		return entity.Referal{}, fmt.Errorf("adminEntity not found")
	}
	s.logger.Info("adminEntity was found, trying to get referal")

	referal, err = s.ReadRepo.GetReferalByRefLink(ctx, referalLink, adminEntity)

	if err != nil {
		return entity.Referal{}, fmt.Errorf("adminEntity not found")
	}
	return referal, nil
}

func (s *ReferalService) generateReferalLink(tgId, inServiceId, adminId int64) string {
	currentTime := time.Now().Format("20060102150405") // Format: YYYYMMDDHHmmss
	return "ref_" + currentTime + "_" + strconv.FormatInt(adminId, 10) + "_" + strconv.FormatInt(tgId, 10) +
		"_" + strconv.FormatInt(inServiceId, 10)
}
