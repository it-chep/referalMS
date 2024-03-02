package admin

import (
	"log/slog"
)

type Admin struct {
	logger *slog.Logger
}

func NewAdmin(logger *slog.Logger) *Admin {
	return &Admin{
		logger: logger,
	}
}
