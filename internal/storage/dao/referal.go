package dao

import "referalMS/internal/domain/entity"

type ReferalDAO struct {
	Id          int64  `sql:"id"`
	TgId        int64  `sql:"tg_id"`
	InServiceId int64  `sql:"in_service_id"`
	AdminId     int64  `sql:"admin_id"`
	Name        string `sql:"name"`
	Username    string `sql:"username"`
	UsersCount  int64  `sql:"users_count"`
}

func NewReferalDAO() *ReferalDAO {
	return &ReferalDAO{}
}

func (dao *ReferalDAO) ToDomain() *entity.Referal {
	return entity.NewReferal(
		dao.TgId,
		dao.AdminId,
		dao.Name,
		entity.WithRefUsername(dao.Username),
		entity.WithRefInServiceId(dao.InServiceId),
		entity.WithRefAllUsersCount(dao.UsersCount),
	)
}
