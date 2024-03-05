package read_repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/internal/storage/dao"
)

type ReferalStorage struct {
	pgClient *pgxpool.Pool
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

func (r *ReferalStorage) GetReferalByRefLink(ctx context.Context, inServiceId int64, referalLink string, admin entity.Admin) (referal entity.Referal, err error) {
	q := `select tg_id, name, username, id_in_integration_service
			from referals r 
				left join admins a on r.admin_id = a.id 
			where r.referal_link = $1 and a.login = $2 and r.id_in_integration_service = $3;
	`
	var referalDAO dao.ReferalDAO
	err = r.pgClient.QueryRow(ctx, q, referalLink, admin.GetLogin(), inServiceId).Scan(&referalDAO)

	if err != nil {
		return entity.Referal{}, err
	}

	referal = *referalDAO.ToDomain()

	return referal, nil
}

func (r *ReferalStorage) GetReferalByTgID(ctx context.Context, tgID, inServiceId int64, admin entity.Admin) (referal entity.Referal, err error) {
	q := `select id, tg_id, name, username 
			from referals r 
				left join admins a on r.admin_id = a.id 
			where r.tg_id = $1 and a.login = $2 and r.id_in_integration_service = $3;
	`
	var referalDAO dao.ReferalDAO
	err = r.pgClient.QueryRow(ctx, q, tgID, admin.GetLogin(), inServiceId).Scan(&referalDAO)
	if err != nil {
		return entity.Referal{}, err
	}

	referal = *referalDAO.ToDomain()

	return referal, nil
}

func (r *ReferalStorage) GetReferalStatistic(ctx context.Context, tgID int64, admin entity.Admin, daysAgo int) (allUsers int64, lastNDays int64, err error) {
	q := `
	with all_users as (
		select count(*) from user u inner join referal r on u.referal_id = r.id
        inner join admins a on a.id = r.admin_id where r.tg_id = $1 and a.login = $2
    ),last_n_days as (
		select count(*) from user u inner join referal r on u.referal_id = r.id 
		inner join admins a on a.id = r.admin_id where r.tg_id = $1 and a.login = $2
		and u.registration_time >= current_date - interval '$3 days'
	)
	select all_users, last_n_days;
	`
	err = r.pgClient.QueryRow(ctx, q, tgID, admin.GetLogin(), daysAgo).Scan(&allUsers, &lastNDays)
	if err != nil {
		return 0, 0, err
	}
	return allUsers, lastNDays, nil
}

func NewReferalStorage(pgClient *pgxpool.Pool, logger *slog.Logger) *ReferalStorage {
	return &ReferalStorage{
		pgClient: pgClient,
		logger:   logger,
	}
}
