package create_user

import (
	"context"
	"log/slog"
	"referalMS/internal/domain/entity"
)

type CreateUserUseCase struct {
	repo   WriteRepo
	logger *slog.Logger
}

func NewCreateUserUseCase(repo WriteRepo, logger *slog.Logger) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, user entity.User, admin entity.Admin) (userId int64, err error) {
	return u.repo.CreateUser(ctx, user, admin)
}
