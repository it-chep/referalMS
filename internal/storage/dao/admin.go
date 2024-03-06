package dao

import (
	"referalMS/internal/domain/entity"
)

type AdminDAO struct {
	id       int64  `sql:"id"`
	login    string `sql:"login"`
	password string `sql:"password"`
	token    string `sql:"token"`
}

func NewAdminDAO(id int64, login, password, token string) *AdminDAO {
	return &AdminDAO{
		id:       id,
		login:    login,
		password: password,
		token:    token,
	}
}

func (dao *AdminDAO) ToDomain() *entity.Admin {
	return entity.NewAdmin(
		dao.login,
		dao.token,
		entity.WithPassword(dao.password),
	)
}
