package create_user

import (
	"context"
	"referalMS/internal/domain/entity"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/mock.go -package mocks . WriteRepo

type WriteRepo interface {
	CreateUser(ctx context.Context, user entity.User, admin entity.Admin) (userID int64, err error)
}
