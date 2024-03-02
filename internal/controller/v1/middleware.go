package v1

import (
	"context"
	"errors"
	"net/http"
	"referalMS/internal/controller/dto"
)

func (api *ApiV1) GetAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		admin, err := api.getAdminHeaders(request)
		if err != nil {
			http.Error(writer, "Some Headers are empty", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(request.Context(), "admin", admin)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func (api *ApiV1) getAdminHeaders(request *http.Request) (admin dto.ExternalAdminDTO, err error) {
	admin = dto.ExternalAdminDTO{
		Login:    request.Header.Get("Admin-Login"),
		Password: request.Header.Get("Admin-Password"),
		Token:    request.Header.Get("Admin-Token"),
	}

	if admin.Login == "" {
		return admin, errors.New("login is required")
	}
	if admin.Password == "" {
		return admin, errors.New("password is required")
	}
	if admin.Token == "" {
		return admin, errors.New("token is required")
	}

	return admin, nil
}
