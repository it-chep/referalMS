package create_user

import (
	"context"
	"referalMS/internal/domain/entity"
)

type CreateUserUseCase struct {
	repo WriteRepo
}

func NewCreateUserUseCase(repo WriteRepo) CreateUserUseCase {
	return CreateUserUseCase{repo: repo}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, referal entity.Referal, user entity.User, admin entity.Admin) (userId int64, err error) {
	return u.repo.CreateUser(ctx, referal, user, admin)
}
