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
	"referalMS/internal/domain/usercases/create_admin"
	"referalMS/internal/domain/usercases/create_referal"
	"referalMS/internal/domain/usercases/create_user"
	"referalMS/internal/storage/read_repo"
	"referalMS/internal/storage/write_repo"
	"referalMS/pkg/client/postgres"
	"time"
)

func (app *App) InitControllers(ctx context.Context) *App {
	app.controller.restController = controller.NewRestController(
		&app.services.adminService, app.services.referalService, app.services.userService, *app.config, app.logger,
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
	app.logger.Info("init pgxclient", app.pgxClient)
	return app
}

func (app *App) InitStorage(ctx context.Context) *App {
	app.storages.adminReadStorage = read_repo.NewAdminStorage(app.pgxClient, app.logger)
	app.storages.adminWriteStorage = write_repo.NewWriteAdminStorage(app.pgxClient, app.logger)
	app.storages.referalWriteStorage = write_repo.NewReferalStorage(app.pgxClient, app.logger)
	app.storages.referalReadStorage = read_repo.NewReferalStorage(app.pgxClient, app.logger)
	app.storages.userWriteStorage = write_repo.NewUserStorage(app.pgxClient, app.logger)
	app.storages.userReadStorage = read_repo.NewUserStorage(app.pgxClient, app.logger)

	app.logger.Info("init admin read storage", app.storages.adminReadStorage)
	app.logger.Info("init admin write storage", app.storages.adminWriteStorage)
	app.logger.Info("init referal read storage", app.storages.referalReadStorage)
	app.logger.Info("init referal write storage", app.storages.referalWriteStorage)
	app.logger.Info("init user read storage", app.storages.userReadStorage)
	app.logger.Info("init user write storage", app.storages.userWriteStorage)

	return app
}

func (app *App) InitUseCases(ctx context.Context) *App {
	app.useCases.createUserUseCase = create_user.NewCreateUserUseCase(app.storages.userWriteStorage, app.logger)
	app.useCases.createReferalUseCase = create_referal.NewCreateReferalUseCase(app.storages.referalWriteStorage, app.logger)
	app.useCases.createAdminUseCase = create_admin.NewCreateAdminUseCase(app.storages.adminWriteStorage, app.logger)
	app.logger.Info("init admin usecase", app.useCases.createAdminUseCase)
	app.logger.Info("init referal usecase", app.useCases.createReferalUseCase)
	app.logger.Info("init user usecase", app.useCases.createUserUseCase)

	return app
}

func (app *App) InitServices(ctx context.Context) *App {
	adminService :=
		admin.NewAdminService(
			app.storages.adminReadStorage,
			app.useCases.createAdminUseCase,
			app.logger,
		)

	app.services = services{
		adminService: adminService,

		userService: user.NewUserService(
			app.useCases.createUserUseCase,
			app.storages.userReadStorage,
			adminService,
			app.logger,
		),
		referalService: referal.NewReferalService(
			app.useCases.createReferalUseCase,
			app.storages.referalReadStorage,
			adminService,
			app.logger,
		),
	}
	app.logger.Info("init services", app.services)
	return app
}
