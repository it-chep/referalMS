package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"referalMS/internal"
	"referalMS/internal/config"
)

func main() {
	ctx := context.Background()
	logger := setupLogger()
	cfg := config.NewConfig()
	log.Fatal(internal.NewApp().Run(ctx, logger, cfg))
}

func setupLogger() *slog.Logger {
	var logger *slog.Logger
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return logger
}
