package v1

import (
	"log/slog"
)

type ApiV1 struct {
	Logger *slog.Logger
}

func NewApiV1(logger *slog.Logger) *ApiV1 {
	return &ApiV1{
		Logger: logger,
	}
}
