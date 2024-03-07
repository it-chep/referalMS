package read_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/internal/storage/dao"
	"referalMS/pkg/client/postgres"
)

type ReferalStorage struct {
	pgClient postgres.Client
	logger   *slog.Logger
}

type Winner struct {
	Name                   string
	TGID                   int64
	Username               string
	IDInIntegrationService int64
	UserCount              int
}

func (r *ReferalStorage) GetReferalByID(ctx context.Context, ID, inServiceId int64, admin entity.Admin) (referal entity.Referal, err error) {
	const op = "internal.storage.read_repo.referal.GetReferalByID"
	q := `select tg_id, name, username, id_in_integration_service
			from referals r 
				left join admins a on r.admin_id = a.id 
			where r.id = $1 and a.login = $2 and r.id_in_integration_service = $3;
	`
	var referalDAO dao.ReferalDAO

	err = r.pgClient.QueryRow(ctx, q, ID, admin.GetLogin(), inServiceId).Scan(&referalDAO)

	if err != nil {
		return entity.Referal{}, err
	}

	referal = *referalDAO.ToDomain()

	return referal, nil
}

func (r *ReferalStorage) GetReferalByRefLink(ctx context.Context, referalLink string, admin entity.Admin) (referal entity.Referal, err error) {
	const op = "internal.storage.read_repo.referal.GetReferalByRefLink"
	q := `select tg_id, name, username, id_in_integration_service
			from referals r 
				left join admins a on r.admin_id = a.id 
			where r.referal_link = $1 and a.login = $2;
	`
	var referalDAO dao.ReferalDAO
	err = r.pgClient.QueryRow(ctx, q, referalLink, admin.GetLogin()).Scan(&referalDAO)

	if err != nil {
		return entity.Referal{}, err
	}

	referal = *referalDAO.ToDomain()

	return referal, nil
}

func (r *ReferalStorage) GetReferalByTgID(ctx context.Context, tgID, inServiceId int64, admin entity.Admin) (referal entity.Referal, err error) {
	const op = "internal.storage.read_repo.referal.GetReferalByTgID"
	q := `select r.id, r.tg_id, r.name, r.username 
			from referals r 
				left join admins a on r.admin_id = a.id 
			where r.tg_id = $1 and a.login = $2 and r.id_in_integration_service = $3;
	`
	var referalDAO dao.ReferalDAO
	err = r.pgClient.QueryRow(ctx, q, tgID, admin.GetLogin(), inServiceId).Scan(
		&referalDAO.Id, &referalDAO.TgId, &referalDAO.Name, &referalDAO.Username,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Referal{}, nil
		}
		return entity.Referal{}, err
	}

	referal = *referalDAO.ToDomain()

	return referal, nil
}

func (r *ReferalStorage) GetReferalStatistic(ctx context.Context, tgID int64, admin entity.Admin, daysAgo int) (allUsers int64, lastNDays int64, err error) {
	const op = "internal.storage.read_repo.referal.GetReferalStatistic"
	r.logger.Info(fmt.Sprintf("op: %s, id: %d", op, tgID))
	q := `
	with all_users as (
		select count(*) from users u inner join referals r on u.referal_id = r.id
        inner join admins a on a.id = r.admin_id where r.tg_id = $1 and a.login = $2
    ),last_n_days as (
		select count(*) from users u inner join referals r on u.referal_id = r.id 
		inner join admins a on a.id = r.admin_id where r.tg_id = $1 and a.login = $2
		and u.registration_time >= current_date - $3 * interval '1 day'
	)
	select (select * from all_users) as all_users, (select * from last_n_days) as last_n_days;
	`
	err = r.pgClient.QueryRow(ctx, q, tgID, admin.GetLogin(), daysAgo).Scan(&allUsers, &lastNDays)
	if err != nil {
		return 0, 0, err
	}
	return allUsers, lastNDays, nil
}

func NewReferalStorage(pgClient postgres.Client, logger *slog.Logger) *ReferalStorage {
	return &ReferalStorage{
		pgClient: pgClient,
		logger:   logger,
	}
}
