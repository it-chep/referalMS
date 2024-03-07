package dao

import (
	"fmt"
	"referalMS/internal/domain/entity"
)

type AdminDAO struct {
	Id       int64  `sql:"id"`
	Login    string `sql:"login"`
	Password string `sql:"password"`
	Token    string `sql:"integrations_token"`
	Salt     int    `sql:"salt"`
}

func NewAdminDAO(id int64, login, password, token string, salt int) *AdminDAO {
	return &AdminDAO{
		Id:       id,
		Login:    login,
		Password: password,
		Token:    token,
		Salt:     salt,
	}
}

func (dao *AdminDAO) ToDomain() *entity.Admin {
	fmt.Println(dao, dao.Id, dao.Login, dao.Password, dao.Password)
	return entity.NewAdmin(
		dao.Login,
		dao.Token,
		entity.WithPassword(dao.Password),
		entity.WithSalt(dao.Salt),
		entity.WithAdmId(dao.Id),
	)
}
