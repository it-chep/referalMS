package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

func (api *ApiV1) CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.CreateUser"
		var resp response.Response
		var user dto.TgUserDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.Logger.Info("Admin:", externalAdminDTO)

		if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
			api.Logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		service := true

		if service {
			resp = response.Ok()
		} else {
			resp = response.Error("Error while create user")
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
