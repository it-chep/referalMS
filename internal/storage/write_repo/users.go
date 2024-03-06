package write_repo

import (
	"context"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/pkg/client/postgres"
)

type WriteUserStorage struct {
	client postgres.Client
	logger *slog.Logger
}

func (u *WriteUserStorage) CreateUser(ctx context.Context, user entity.User, admin entity.Admin) (userID int64, err error) {
	q := `
	with r_id as (
		select id from referal where referal_link = $3
	), a_id as (
		select id from admins where login = $1
	)
	insert into user (admin_id, tg_id, referal_link, username, in_service_id ,referal_id) 
	values (a_id, $2, $3, $4, $5, r_id) returning id;
	`
	err = u.client.QueryRow(
		ctx, q, admin.GetLogin(), user.GetTgId(), user.GetReferalLink(), user.GetUsername(), user.GetInServiceId(),
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func NewUserStorage(client postgres.Client, logger *slog.Logger) *WriteUserStorage {
	return &WriteUserStorage{
		client: client,
		logger: logger,
	}
}
