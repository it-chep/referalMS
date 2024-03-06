package admin

import (
	"log/slog"
)

type Admin struct {
	adminService AdminService
	logger       *slog.Logger
}

func NewAdmin(adminService AdminService, logger *slog.Logger) *Admin {
	return &Admin{
		adminService: adminService,
		logger:       logger,
	}
}
