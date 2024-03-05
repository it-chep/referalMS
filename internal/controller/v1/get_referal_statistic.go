package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

func (api *ApiV1) GetReferalStatistic(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.GetReferalStatistic"
		var referalUser dto.ReferalStatisticDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.logger.Info("Admin:", externalAdminDTO, op)

		if err := json.NewDecoder(request.Body).Decode(&referalUser); err != nil {
			api.logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		allUsers, usersLastNDays, err := api.referalService.GetReferalStatistic(ctx, referalUser, externalAdminDTO)
		if err != nil {
			api.logger.Error(fmt.Sprintf("Error while get statistig for referal. TgId: %d, error: %s, op: %s", referalUser.TgId, err, op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := dto.GetReferalStatisticDTO{
			Status:    response.StatusOk,
			AllUsers:  allUsers,
			LastNDays: usersLastNDays,
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
