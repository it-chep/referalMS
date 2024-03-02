package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

type GetReferalStatisticResponse struct {
	Status    string `json:"status"`
	Error     string `json:"error"`
	AllUsers  int64  `json:"all_users"`
	LastNDays int64  `json:"last_n_days"`
}

func (api *ApiV1) GetReferalStatistic() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.GetReferalStatistic"
		var referalUser dto.ReferalStatisticDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.Logger.Info("Admin:", externalAdminDTO)

		if err := json.NewDecoder(request.Body).Decode(&referalUser); err != nil {
			api.Logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		// Service logic

		status := response.StatusOk
		resp := GetReferalStatisticResponse{
			Status:    status,
			AllUsers:  0,
			LastNDays: 0,
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
