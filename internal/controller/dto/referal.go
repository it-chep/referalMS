package dto

// ReferalUserDTO representation of referal
type ReferalUserDTO struct {
	TgId        int64  `json:"tg_id"`
	InServiceID int64  `json:"in_service_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
}

// ReferalStatisticDTO referal tg_id
type ReferalStatisticDTO struct {
	TgId int64 `json:"tg_id"`
}

// GetWinnersDTO get top referal DTO
type GetWinnersDTO struct {
	Limit        int `json:"limit"`
	DaysInterval int `json:"days_interval"`
}

type GetReferalStatisticDTO struct {
	Status    string `json:"status"`
	Error     string `json:"error"`
	AllUsers  int64  `json:"all_users"`
	LastNDays int64  `json:"last_n_days"`
}
