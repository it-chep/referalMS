package create_referal

import "context"

type CreateRferalUseCase struct {
	repo WriteRepo
}

func NewCreateRferalUseCase(repo WriteRepo) CreateRferalUseCase {
	return CreateRferalUseCase{repo: repo}
}

func (uc CreateRferalUseCase) Create(ctx context.Context) error {
	return uc.repo.Create(ctx)
}
