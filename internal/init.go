package internal

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"referalMS/internal/controller"
	"referalMS/internal/domain/service/admin"
	"referalMS/internal/domain/service/referal"
	"referalMS/internal/domain/service/user"
	"referalMS/internal/domain/usercases/create_referal"
	"referalMS/internal/domain/usercases/create_user"
	"referalMS/internal/storage/read_repo"
	"referalMS/internal/storage/write_repo"
	"referalMS/pkg/client/postgres"
	"time"
)

func (app *App) InitControllers(ctx context.Context) *App {
	app.controller.restController = controller.NewRestController(
		&app.services.adminService, &app.services.referalService, app.services.userService, *app.config, app.logger,
	)
	app.controller.restController.InitController(ctx)

	app.server = &http.Server{
		Addr:         app.config.HTTPServer.Address,
		Handler:      app.controller.restController,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}
	return app
}

func (app *App) InitLogger(ctx context.Context) *App {
	app.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return app
}

func (app *App) InitPgxConn(ctx context.Context) *App {
	client, err := postgres.NewClient(ctx, app.config.StorageConfig)
	if err != nil {
		log.Fatal(err)
	}
	app.pgxClient = client
	return app
}

func (app *App) InitStorage(ctx context.Context) *App {
	app.storages.adminReadStorage = read_repo.NewAdminStorage(app.pgxClient, app.logger)
	app.storages.referalWriteStorage = *write_repo.NewReferalStorage(app.pgxClient, app.logger)
	app.storages.referalReadStorage = read_repo.NewReferalStorage(app.pgxClient, app.logger)
	app.storages.userWriteStorage = *write_repo.NewUserStorage(app.pgxClient, app.logger)
	app.storages.userReadStorage = read_repo.NewUserStorage(app.pgxClient, app.logger)
	return app
}

func (app *App) InitUseCases(ctx context.Context) *App {
	app.useCases.createUserUseCase = create_user.NewCreateUserUseCase(&app.storages.userWriteStorage)
	app.useCases.createReferalUseCase = create_referal.NewCreateReferalUseCase(&app.storages.referalWriteStorage)
	return app
}

func (app *App) InitServices(ctx context.Context) *App {
	app.services = services{
		userService: user.NewUserService(
			app.useCases.createUserUseCase,
			app.storages.userReadStorage,
			app.logger,
		),
		referalService: referal.NewReferalService(
			app.useCases.createReferalUseCase,
			app.storages.referalReadStorage,
			app.services.adminService,
			app.logger,
		),
		adminService: admin.NewAdminService(
			app.storages.adminReadStorage
		),
	}
	return app
}
