package internal

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log/slog"
	"net/http"
	"referalMS/internal/config"
	"referalMS/internal/controller"
	"referalMS/internal/domain/service/admin"
	"referalMS/internal/domain/service/referal"
	"referalMS/internal/domain/service/user"
	"referalMS/internal/storage/write_repo"
)

type controllers struct {
	restController *controller.RestController
}

type services struct {
	userService    user.UserService
	referalService referal.ReferalService
	adminService   admin.AdminService
}

type useCases struct {
	createReferalUseCase
	createUserUseCase
}

type storages struct {
	adminReadStorage    admin.ReadAdminStorage
	referalReadStorage  referal.ReadReferalStorage
	referalWriteStorage write_repo.WriteReferalStorage
	userReadStorage     user.ReadUserStorage
	userWriteStorage    write_repo.WriteUserStorage
}

type App struct {
	logger     *slog.Logger
	config     *config.Config
	controller controllers
	services   services
	storages   storages
	useCases   useCases
	pgxClient  *pgxpool.Pool
	server     *http.Server
}

func NewApp(ctx context.Context) *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.InitLogger(ctx).
		InitPgxConn(ctx).
		InitUseCases(ctx).
		InitServices(ctx).
		InitControllers(ctx)

	return app
}

func (app *App) Run() error {
	app.logger.Info("start server")
	return app.server.ListenAndServe()
}
