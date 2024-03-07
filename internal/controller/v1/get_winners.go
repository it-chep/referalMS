package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

func (api *ApiV1) GetWinners(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.GetWinners"
		var winnersFilter dto.GetWinnersDTO
		var users []dto.WinnerReferal

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.logger.Info("Admin:", externalAdminDTO, op)

		err := json.NewDecoder(request.Body).Decode(&winnersFilter)
		if err != nil {
			api.logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		if winnersFilter.Limit == 0 {
			api.logger.Warn(fmt.Sprintf("Limit winners is empty %s", op))
			winnersFilter.Limit = api.cfg.ReferalConfig.Winners.Limit
		}

		if winnersFilter.DaysInterval == 0 {
			api.logger.Warn(fmt.Sprintf("Days interval winners is empty %s", op))
			winnersFilter.DaysInterval = api.cfg.ReferalConfig.Winners.Interval
		}

		winners, err := api.adminService.GetWinners(ctx, externalAdminDTO, winnersFilter)

		for _, winner := range winners {
			winnerDTO := dto.WinnerReferal{
				Name:        winner.GetName(),
				UsersCount:  winner.GetAllUsers(),
				Username:    winner.GetUsername(),
				InServiceId: winner.GetInServiceId(),
			}
			users = append(users, winnerDTO)
		}

		resp := dto.GetWinnersResponse{
			Status: response.StatusOk,
			Users:  users,
		}

		if err != nil {
			api.logger.Error(fmt.Sprintf("Error while get winners. Error: %s OP: %s", err, op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(resp)
		if err != nil {
			api.logger.Error(fmt.Sprintf("Error while serializing JSON %s", op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		_, err = writer.Write(jsonData)
		if err != nil {
			api.logger.Error(fmt.Sprintf("SERVER ERROR %s, op: %s", err, op))
			return
		}
	}
}
