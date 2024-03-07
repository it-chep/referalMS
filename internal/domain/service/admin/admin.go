package admin

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"math/rand"
	"referalMS/internal/controller/dto"
	"referalMS/internal/domain/entity"
	"strconv"
	"time"
)

const weightOfHashing = 10

type AdminService struct {
	readRepo           ReadAdminStorage
	createAdminUseCase CreateAdminUseCase
	logger             *slog.Logger
}

func NewAdminService(readRepo ReadAdminStorage, createAdminUseCase CreateAdminUseCase, logger *slog.Logger) AdminService {
	return AdminService{
		readRepo:           readRepo,
		createAdminUseCase: createAdminUseCase,
		logger:             logger,
	}
}

func (s *AdminService) GetAdmin(ctx context.Context, dto dto.ExternalAdminDTO) (admin entity.Admin, err error) {
	const op = "internal.domain.service.admin.admin.GetAdmin"
	admin, err = s.readRepo.GetAdmin(
		ctx,
		dto.Login,
		dto.Token,
	)

	if err != nil {
		return entity.Admin{}, err
	}

	err = s.checkPassEqualPassHash(dto.Password, admin.GetSalt(), admin.GetPassword())

	if err != nil {
		return entity.Admin{}, err
	}

	return admin, nil
}

func (s *AdminService) GetWinners(
	ctx context.Context, dto dto.ExternalAdminDTO, filters dto.GetWinnersDTO,
) (winners []entity.Referal, err error) {
	admin, err := s.GetAdmin(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("admin not found")
	}

	winnersFilter := entity.NewWinnersFilter(
		filters.Limit,
		filters.DaysInterval,
	)

	winners, err = s.readRepo.GetWinners(ctx, admin, *winnersFilter)
	if err != nil {
		return nil, err
	}

	return winners, nil
}

func (s *AdminService) RegisterNewAdmin(ctx context.Context, adminDTO dto.AdminDTO) (adminId int64, err error) {
	adminFromRepo, err := s.GetAdmin(ctx, dto.ExternalAdminDTO(adminDTO))
	if adminFromRepo != (entity.Admin{}) {
		s.logger.Info(fmt.Sprintf("admin with login %s exist", adminDTO.Login))
		return -1, err
	}

	hashedPassword, salt, err := s.generatePasswordHash(adminDTO.Password)

	if err != nil {
		return -1, err
	}
	admin := *entity.NewAdmin(adminDTO.Login, adminDTO.Token, entity.WithPassword(hashedPassword), entity.WithSalt(salt))

	s.logger.Info("admin Entity", admin)
	adminId, err = s.createAdminUseCase.Execute(ctx, admin)

	if err != nil {
		return -1, err
	}
	return adminId, nil
}

func (s *AdminService) generatePasswordHash(password string) (string, int, error) {
	rand.Seed(time.Now().UnixNano())
	salt := rand.Intn(10000)
	hash, err := bcrypt.GenerateFromPassword([]byte(password+strconv.Itoa(salt)), weightOfHashing)

	return string(hash), salt, err
}

func (s *AdminService) checkPassEqualPassHash(password string, salt int, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password+strconv.Itoa(salt)))
	return err
}
