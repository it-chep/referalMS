package create_referal

import "context"

//go:generate mockgen -source interfaces.go -destination ./mocks/mock.go -package mocks . WriteRepo

type WriteRepo interface {
	Create(ctx context.Context) error
}
