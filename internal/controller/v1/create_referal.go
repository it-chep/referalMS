package v1

import (
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

func (api *ApiV1) CreateReferal() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.CreateReferal"

		var referalUser dto.ReferalUserDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.Logger.Info("Admin:", externalAdminDTO)

		if err := json.NewDecoder(request.Body).Decode(&referalUser); err != nil {
			api.Logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		// Service logic

		status := response.StatusOk
		resp := CreateReferalResponse{
			Status:      status,
			ReferalLink: "SomeString",
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
