package dao

import "referalMS/internal/domain/entity"

type ReferalDAO struct {
	tgId        int64  `sql:"tg_id"`
	inServiceId int64  `sql:"in_service_id"`
	adminId     int64  `sql:"admin_id"`
	name        string `sql:"name"`
	username    string `sql:"username"`
	usersCount  string
}

func NewReferalDAO() *ReferalDAO {
	return &ReferalDAO{}
}

func (dao *ReferalDAO) ToDomain() *entity.Referal {
	return entity.NewReferal(
		dao.tgId,
		dao.adminId,
		dao.name,
		entity.WithRefUsername(dao.username),
		entity.WithRefInServiceId(dao.inServiceId),
	)
}
