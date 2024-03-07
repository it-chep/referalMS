package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"referalMS/internal/controller/dto"
	"referalMS/pkg/api/response"
)

func (admn *Admin) CreateAdmin(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "controller.controller.v1.CreateUser"
		var adminDTO dto.AdminDTO

		if err := json.NewDecoder(request.Body).Decode(&adminDTO); err != nil {
			admn.logger.Error(fmt.Sprintf("Error while parsing JSON %s", op))
			http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		_, err := admn.adminService.RegisterNewAdmin(ctx, adminDTO)
		if err != nil {
			admn.logger.Error(fmt.Sprintf("Error while creating admin err: %s op: %s", err, op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := response.Ok()

		jsonData, err := json.Marshal(resp)
		if err != nil {
			admn.logger.Error(fmt.Sprintf("Error while serializing JSON %s", op))
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		_, err = writer.Write(jsonData)
		if err != nil {
			admn.logger.Error(fmt.Sprintf("SERVER ERROR %s, op: %s", err, op))
			return
		}
	}
}
