package write_repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log/slog"
	"referalMS/internal/domain/entity"
)

type WriteReferalStorage struct {
	pgClient *pgxpool.Pool
	logger   *slog.Logger
}

func (r *WriteReferalStorage) CreateReferal(ctx context.Context, referal entity.Referal, admin entity.Admin) (referalId int64, err error) {
	q := `
	with a_id as ( select id from admins where login = $1)
	insert into referal
	(admin_id, tg_id, id_in_integration_service, name, username, referal_link) 
	values (a_id, $2, $3, $4, $5, $6)
	`
	err = r.pgClient.QueryRow(
		ctx, q, admin.GetLogin(), referal.GetTgId(), referal.GetInServiceId(),
		referal.GetName(), referal.GetUsername(), referal.GetReferalLink(),
	).Scan(&referalId)
	if err != nil {
		return 0, err
	}
	return referalId, nil
}

func NewReferalStorage(pgClient *pgxpool.Pool, logger *slog.Logger) *WriteReferalStorage {
	return &WriteReferalStorage{
		pgClient: pgClient,
		logger:   logger,
	}
}
