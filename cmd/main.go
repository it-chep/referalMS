package main

import (
	"log/slog"
	"os"
)

const (
	envLocal = "test"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	//	TODO init config
	logger := setupLogger(envLocal)

	logger.Info("start app")

	router := NewRouter()
	router.InitRouter()
	// Auth middleware
	//	TODO init repo
	//	TODO init service

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	//switch env {
	//
	//}
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
