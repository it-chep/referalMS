package dto

import "referalMS/internal/domain/entity"

// ExternalAdminDTO get Headers for integration
type ExternalAdminDTO struct {
	Login    string
	Password string
	Token    string
}

// AdminDTO
type AdminDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type WinnerReferal struct {
	Name        string `json:"name"`
	UsersCount  int64  `json:"users_count"`
	Username    string `json:"username"`
	InServiceId int64  `json:"in_service_id"`
}

type GetWinnersResponse struct {
	Status string          `json:"status"`
	Error  string          `json:"error"`
	Users  []WinnerReferal `json:"users"`
}

func (a *ExternalAdminDTO) ToDomain(password string) *entity.Admin {
	return entity.NewAdmin(a.Login, a.Token, entity.WithPassword(password))
}
