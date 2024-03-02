package dto

// TgUserDTO ...
type TgUserDTO struct {
	TgId        int64  `json:"tg_id"`
	InServiceID int64  `json:"in_service_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	ReferalLink string `json:"referal_link"`
}
