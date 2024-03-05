package dao

import "referalMS/internal/domain/entity"

type UserDAO struct {
	tgId        int64  `sql:"tg_id"`
	adminId     int64  `sql:"admin_id"`
	inServiceId int64  `sql:"in_service_id"`
	referalId   int64  `sql:"referal_id"`
	name        string `sql:"name"`
	username    string `sql:"username"`
	referalLink string `sql:"referal_link"`
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) ToDomain() *entity.User {
	return entity.NewUser(
		dao.tgId,
		dao.adminId,
		dao.name,
		//entity.WithUsrReferalId(dao.referalId),
	)
}
