package internal

import (
	"context"
	"log/slog"
	"net/http"
	"referalMS/internal/config"
	"referalMS/internal/controller"
	"referalMS/internal/domain/service/referals"
	"referalMS/internal/domain/service/users"
	"time"
)

type controllers struct {
	restController *controller.RestController
}

type services struct {
	userService    users.UserService
	referalService referals.ReferalService
}

type App struct {
	controller controllers
	services   services
	server     *http.Server
}

func NewApp() *App {
	app := &App{}

	app.InitPgxConn().
		InitControllers()

	return app
}

func (app *App) Run(ctx context.Context, logger *slog.Logger, cfg *config.Config) error {
	control := controller.NewRestController(logger)
	control.InitController(ctx)

	logger.Info("start server")

	app.server = &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      control,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}
	return app.server.ListenAndServe()
}
