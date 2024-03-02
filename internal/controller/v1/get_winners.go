package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

type WinnerReferal struct {
	UsersCount  int64  `json:"users_count"`
	Username    string `json:"username"`
	InServiceId int64  `json:"in_service_id"`
}

type GetWinnersResponse struct {
	Status string          `json:"status"`
	Error  string          `json:"error"`
	Users  []WinnerReferal `json:"users"`
}

func (api *ApiV1) GetWinners(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.GetWinners"
		var winnersFilter dto.GetWinnersDTO
		var users []WinnerReferal

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.Logger.Info("Admin:", externalAdminDTO)

		err := json.NewDecoder(request.Body).Decode(&winnersFilter)
		if err != nil {
			api.Logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		// TODO GET ADMIN

		if winnersFilter.Limit == 0 {
			api.Logger.Warn(fmt.Sprintf("Limit winners is empty %s", op))
			winnersFilter.Limit = 3
			//admin.Limit = cfg.ReferalConfig.Winners.Limit
		}

		if winnersFilter.DaysInterval == 0 {
			api.Logger.Warn(fmt.Sprintf("Days interval winners is empty %s", op))
			//admin.DaysInterval = cfg.ReferalConfig.Winners.Interval
		}

		//
		for i := 1; i <= winnersFilter.Limit; i++ {
			user := WinnerReferal{
				UsersCount:  int64(i * 2),
				Username:    fmt.Sprintf("user%d", i),
				InServiceId: int64(i * 123),
			}
			users = append(users, user)
		}
		//
		status := response.StatusOk
		resp := GetWinnersResponse{
			Status: status,
			Users:  users,
		}

		jsonData, err := json.Marshal(resp)
		if err != nil {
			api.Logger.Error(fmt.Sprintf("Error while serializing JSON %s", op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		writer.Write(jsonData)
	}
}
