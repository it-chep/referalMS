package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

func (api *ApiV1) CreateUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.CreateUser"
		var resp response.Response
		var user dto.TgUserDTO

		externalAdminDTO, _ := request.Context().Value("admin").(dto.ExternalAdminDTO)
		api.logger.Info("Admin:", externalAdminDTO, op)

		if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
			api.logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		api.userService.RegisterNewUser(ctx)

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
