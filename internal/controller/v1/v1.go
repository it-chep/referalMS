package v1

import (
	"log/slog"
	"referalMS/internal/config"
	"referalMS/internal/controller"
)

type ApiV1 struct {
	adminService   controller.AdminService
	referalService controller.ReferalService
	userService    controller.UserService
	cfg            config.Config
	logger         *slog.Logger
}

func NewApiV1(
	adminService controller.AdminService,
	referalService controller.ReferalService,
	userService controller.UserService,
	cfg config.Config,
	logger *slog.Logger,
) *ApiV1 {

	return &ApiV1{
		adminService:   adminService,
		referalService: referalService,
		userService:    userService,
		cfg:            cfg,
		logger:         logger,
	}

}
