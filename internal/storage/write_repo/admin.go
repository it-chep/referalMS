package write_repo

import (
	"context"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/pkg/client/postgres"
	"time"
)

type WriteAdminStorage struct {
	pgClient postgres.Client
	logger   *slog.Logger
}

func NewWriteAdminStorage(pgClient postgres.Client, logger *slog.Logger) *WriteAdminStorage {
	return &WriteAdminStorage{
		pgClient: pgClient,
		logger:   logger,
	}
}

func (a *WriteAdminStorage) CreateAdmin(ctx context.Context, admin entity.Admin) (adminId int64, err error) {
	q := `
	insert into admins (login, password, integrations_token, integrator_id, last_login_time, registration_time, salt) 
	values ($1, $2, $3, 1, $4, $5, $6) returning id;
`
	err = a.pgClient.QueryRow(
		ctx, q, admin.GetLogin(), admin.GetPassword(), admin.GetIntegrationsToken(), admin.GetLastLogin(), time.Now(), admin.GetSalt(),
	).Scan(&adminId)
	if err != nil {
		return 0, err
	}
	return adminId, nil
}
