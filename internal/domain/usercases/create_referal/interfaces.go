package create_referal

import (
	"context"
	"referalMS/internal/domain/entity"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/mock.go -package mocks . WriteRepo

type WriteRepo interface {
	CreateReferal(ctx context.Context, referal entity.Referal, admin entity.Admin) (referalId int64, err error)
}
