package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

type CreateReferalResponse struct {
	Status      string `json:"status"`
	Error       string `json:"error"`
	ReferalLink string `json:"referal_link"`
}

func (api *ApiV1) CreateReferal(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.CreateReferal"
		var resp CreateReferalResponse
		var referalUser dto.ReferalUserDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.logger.Info("Admin:", externalAdminDTO, op)

		if err := json.NewDecoder(request.Body).Decode(&referalUser); err != nil {
			api.logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		referalLink, err := api.referalService.RegisterNewReferal(ctx, referalUser, externalAdminDTO)
		if err != nil {
			api.logger.Error(fmt.Sprintf("error while creating referal. Error: %s OP: %s", err, op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp = CreateReferalResponse{
			Status:      response.StatusOk,
			ReferalLink: referalLink,
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
