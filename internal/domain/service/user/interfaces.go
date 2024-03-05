package user

import (
	"context"
	"referalMS/internal/domain/entity"
)

type ReadUserStorage interface {
	GetUserByID(ctx context.Context, userID int64) (user entity.User, err error)
	GetUserByTgID(ctx context.Context, tgID int64) (user entity.User, err error)
	GetUserByReferalLink(ctx context.Context, referalLink string) (user entity.User, err error)
	GetUserByReferalId(ctx context.Context, referalId int64) (user entity.User, err error)
	GetAllUsers(ctx context.Context) (users []entity.User, err error)
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, referal entity.Referal, user entity.User, admin entity.Admin) (userId int64, err error)
}
