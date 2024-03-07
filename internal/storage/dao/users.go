package dao

import "referalMS/internal/domain/entity"

type UserDAO struct {
	TgId        int64  `sql:"tg_id"`
	AdminId     int64  `sql:"admin_id"`
	InServiceId int64  `sql:"in_service_id"`
	ReferalId   int64  `sql:"referal_id"`
	Name        string `sql:"name"`
	Username    string `sql:"username"`
	ReferalLink string `sql:"referal_link"`
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) ToDomain() *entity.User {
	return entity.NewUser(
		dao.TgId,
		dao.AdminId,
		dao.Name,
	)
}
