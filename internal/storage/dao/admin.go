package dao

import (
	"referalMS/internal/domain/entity"
)

type AdminDAO struct {
	login    string `sql:"login"`
	password string `sql:"password"`
	token    string `sql:"token"`
}

func NewAdminDAO(login, password, token string) *AdminDAO {
	return &AdminDAO{
		login,
		password,
		token,
	}
}

func (dao *AdminDAO) ToDomain() *entity.Admin {
	return entity.NewAdmin(
		dao.login,
		dao.password,
		dao.token,
	)
}
