package create_admin

import (
	"context"
	"referalMS/internal/domain/entity"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/mock.go -package mocks . WriteRepo

type WriteRepo interface {
	CreateAdmin(ctx context.Context, admin entity.Admin) (adminId int64, err error)
}
