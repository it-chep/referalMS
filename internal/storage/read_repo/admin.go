package read_repo

import (
	"context"
	"fmt"
	"log/slog"
	"referalMS/internal/domain/entity"
	"referalMS/internal/storage/dao"
	"referalMS/pkg/client/postgres"
)

type AdminStorage struct {
	pgClient postgres.Client
	logger   *slog.Logger
}

func NewAdminStorage(pgClient postgres.Client, logger *slog.Logger) *AdminStorage {
	return &AdminStorage{
		pgClient: pgClient,
		logger:   logger,
	}
}

func (a *AdminStorage) GetAdmin(ctx context.Context, login, token string) (admin entity.Admin, err error) {
	q := `select id, login, password, integrations_token, salt from admins where login = $1 and integrations_token = $2;`

	var adminDAO dao.AdminDAO
	err = a.pgClient.QueryRow(ctx, q, login, token).Scan(&adminDAO.Id, &adminDAO.Login, &adminDAO.Password, &adminDAO.Token, &adminDAO.Salt)
	admin = *adminDAO.ToDomain()
	a.logger.Info("GetAdmin:", admin, adminDAO.Id)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Err: %s", err))
		return entity.Admin{}, err
	}
	return admin, nil
}

func (a *AdminStorage) GetWinners(ctx context.Context, admin entity.Admin, winnersFilter entity.WinnersFilter) (winners []entity.Referal, err error) {
	q := `
		select r.name, r.tg_id, r.username, r.id_in_integration_service, COUNT(u.id) as user_count 
		from referals r 
		    inner join users u on r.id = u.referal_id 
		    inner join admins a on u.admin_id = a.id
     	where a.login = $1 and u.registration_time >= current_date - $2 * interval '1 day'
     		group by r.id, r.name, r.tg_id, r.username, r.id_in_integration_service
     		order by user_count desc
     	limit $3;
	`

	rows, err := a.pgClient.Query(ctx, q, admin.GetLogin(), winnersFilter.GetDaysInterval(), winnersFilter.GetLimit())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	winners = make([]entity.Referal, 0)

	for rows.Next() {
		var referalDAO dao.ReferalDAO

		err := rows.Scan(&referalDAO.Name, &referalDAO.TgId, &referalDAO.Username, &referalDAO.InServiceId, &referalDAO.UsersCount)
		if err != nil {
			return nil, err
		}

		refEntity := referalDAO.ToDomain()

		winners = append(winners, *refEntity)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return winners, nil
}
