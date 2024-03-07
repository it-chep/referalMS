package read_repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/internal/storage/dao"
	"referalMS/pkg/client/postgres"
)

type UserStorage struct {
	client postgres.Client
	logger *slog.Logger
}

func (u *UserStorage) GetUserByID(ctx context.Context, userID int64) (user entity.User, err error) {
	q := "select tg_id, admin_id, username, referal_link from users where id = $1"

	var userDAO dao.UserDAO

	err = u.client.QueryRow(ctx, q, userID).Scan(&userDAO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Create and return a new User entity
	user = *userDAO.ToDomain()
	return user, nil
}

func (u *UserStorage) GetUserByTgID(ctx context.Context, tgID int64) (user entity.User, err error) {
	q := "select admin_id, username, referal_link from users where tg_id = $1"
	var userDAO dao.UserDAO

	err = u.client.QueryRow(ctx, q, tgID).Scan(&userDAO.AdminId, &userDAO.Username, &userDAO.ReferalLink)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Create and return a new User entity
	user = *userDAO.ToDomain()
	return user, nil
}

func (u *UserStorage) GetUserByReferalLink(ctx context.Context, referalLink string) (user entity.User, err error) {
	q := "select tg_id, admin_id, username, referal_link from users where referal_link = $1"
	var userDAO dao.UserDAO

	err = u.client.QueryRow(ctx, q, referalLink).Scan(&userDAO.TgId, &userDAO.AdminId, &userDAO.Username, &userDAO.ReferalLink)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Create and return a new User entity
	user = *userDAO.ToDomain()
	return user, nil
}

func (u *UserStorage) GetUserByReferalId(ctx context.Context, referalId int64) (user entity.User, err error) {
	q := "select tg_id, admin_id, username, referal_link from users where referal_id = $1"
	var userDAO dao.UserDAO

	err = u.client.QueryRow(ctx, q, referalId).Scan(&userDAO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Create and return a new User entity
	user = *userDAO.ToDomain()
	return user, nil
}

func (u *UserStorage) GetAllUsers(ctx context.Context) (users []entity.User, err error) {
	q := "select tg_id, admin_id, username, referal_link from users"
	rows, err := u.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userDAO dao.UserDAO
		err := rows.Scan(&userDAO)
		if err != nil {
			return nil, err
		}
		user := *userDAO.ToDomain()
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserStorage(client postgres.Client, logger *slog.Logger) *UserStorage {
	return &UserStorage{
		client: client,
		logger: logger,
	}
}
