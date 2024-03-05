package dto

// ExternalAdminDTO get Headers for integration
type ExternalAdminDTO struct {
	Login    string
	Password string
	Token    string
}

type WinnerReferal struct {
	UsersCount  int64  `json:"users_count"`
	Username    string `json:"username"`
	InServiceId int64  `json:"in_service_id"`
}

type GetWinnersResponse struct {
	Status string          `json:"status"`
	Error  string          `json:"error"`
	Users  []WinnerReferal `json:"user"`
}
