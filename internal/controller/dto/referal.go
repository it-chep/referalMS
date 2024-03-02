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

// GetWinnersDTO get top referals DTO
type GetWinnersDTO struct {
	Limit        int `json:"limit"`
	DaysInterval int `json:"days_interval"`
}
