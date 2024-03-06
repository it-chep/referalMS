package admin

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"math/rand"
	"os"
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
	admin, err = s.readRepo.GetAdmin(
		ctx,
		dto.Login,
		dto.Token,
	)
	if err != nil {
		return entity.Admin{}, err
	}
	err = s.checkPassEqualPassHash(dto.Password, admin.GetPassword())
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
	hashedPassword, err := s.generatePasswordHash(adminDTO.Password)
	s.logger.Info("Hashed password", hashedPassword)
	if err != nil {
		return -1, err
	}
	admin := *entity.NewAdmin(adminDTO.Login, adminDTO.Token, entity.WithPassword(hashedPassword))
	s.logger.Info("admin Entity", admin)
	adminId, err = s.createAdminUseCase.Execute(ctx, admin)
	if err != nil {
		return -1, err
	}
	return adminId, nil
}

func (s *AdminService) generatePasswordHash(password string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	salt := rand.Intn(10000)
	hash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("PASS_SALT")+password+strconv.Itoa(salt)), weightOfHashing)

	return string(hash), err
}

func (s *AdminService) checkPassEqualPassHash(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}
