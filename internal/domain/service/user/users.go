package user

import (
	"context"
	"log/slog"
)

type UserService struct {
	createUserUseCase CreateUserUseCase
	readRepo          ReadUserStorage
	logger            *slog.Logger
}

func NewUserService(CreateUserUseCase CreateUserUseCase, readRepo ReadUserStorage, logger *slog.Logger) UserService {
	return UserService{
		createUserUseCase: CreateUserUseCase,
		readRepo:          readRepo,
		logger:            logger,
	}
}

func (u UserService) RegisterNewUser(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
