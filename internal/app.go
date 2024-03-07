package internal

import (
	"context"
	"log/slog"
	"net/http"
	"referalMS/internal/config"
	"referalMS/internal/controller"
	v1 "referalMS/internal/controller/v1"
	"referalMS/internal/domain/service/admin"
	"referalMS/internal/domain/service/referal"
	"referalMS/internal/domain/service/user"
	"referalMS/internal/domain/usercases/create_admin"
	"referalMS/internal/domain/usercases/create_referal"
	"referalMS/internal/domain/usercases/create_user"
	"referalMS/pkg/client/postgres"
)

type controllers struct {
	restController *controller.RestController
}

type services struct {
	userService    v1.UserService
	referalService v1.ReferalService
	adminService   admin.AdminService
}

type useCases struct {
	createReferalUseCase referal.CreateReferalUseCase
	createUserUseCase    user.CreateUserUseCase
	createAdminUseCase   admin.CreateAdminUseCase
}

type storages struct {
	adminReadStorage    admin.ReadAdminStorage
	adminWriteStorage   create_admin.WriteRepo
	referalReadStorage  referal.ReadReferalStorage
	referalWriteStorage create_referal.WriteRepo
	userReadStorage     user.ReadUserStorage
	userWriteStorage    create_user.WriteRepo
}

type App struct {
	logger     *slog.Logger
	config     *config.Config
	controller controllers
	services   services
	storages   storages
	useCases   useCases
	pgxClient  postgres.Client
	server     *http.Server
}

func NewApp(ctx context.Context) *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.InitLogger(ctx).
		InitPgxConn(ctx).
		InitStorage(ctx).
		InitUseCases(ctx).
		InitServices(ctx).
		InitControllers(ctx)

	return app
}

func (app *App) Run() error {
	app.logger.Info("start server")
	return app.server.ListenAndServe()
}
