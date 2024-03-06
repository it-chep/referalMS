package dto

import "referalMS/internal/domain/entity"

// TgUserDTO ...
type TgUserDTO struct {
	TgId        int64  `json:"tg_id"`
	InServiceID int64  `json:"in_service_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	ReferalLink string `json:"referal_link"`
}

func (u *TgUserDTO) ToDomain(adminId int64) entity.User {
	return *entity.NewUser(
		u.TgId,
		adminId,
		u.Name,
		entity.WithUsrInServiceId(u.InServiceID),
		entity.WithUsrReferalLink(u.ReferalLink),
		entity.WithUsrUsername(u.Username),
	)
}
