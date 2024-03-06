package controller

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"referalMS/internal/config"
	"referalMS/internal/controller/admin"
	v1 "referalMS/internal/controller/v1"
)

type RestController struct {
	router         *chi.Mux
	adminService   admin.AdminService
	referalService v1.ReferalService
	userService    v1.UserService
	cfg            config.Config
	logger         *slog.Logger
}

func NewRestController(
	adminService admin.AdminService,
	referalService v1.ReferalService,
	userService v1.UserService,
	cfg config.Config,
	logger *slog.Logger,
) *RestController {

	const op = "controller.controller.NewRouter"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	logger.Info(fmt.Sprintf("init router %s", op))

	return &RestController{
		router:         r,
		adminService:   adminService,
		referalService: referalService,
		userService:    userService,
		cfg:            cfg,
		logger:         logger,
	}
}

func (r *RestController) InitController(ctx context.Context) {
	const op = "controller.controller.InitController"
	apiV1Router := chi.NewRouter()
	apiV1 := v1.NewApiV1(r.adminService, r.referalService, r.userService, r.cfg, r.logger)
	apiV1Router.Use(apiV1.GetAdminMiddleware)

	admn := admin.NewAdmin(r.adminService, r.logger)

	apiV1Router.Post("/new_referal/", apiV1.CreateReferal(ctx))
	apiV1Router.Post("/new_user/", apiV1.CreateUser(ctx))
	apiV1Router.Post("/get_statistic/", apiV1.GetReferalStatistic(ctx))
	apiV1Router.Post("/get_winners/", apiV1.GetWinners(ctx))

	r.router.Route("/api", func(r chi.Router) {
		r.Mount("/v1", apiV1Router)
	})

	r.router.Route("/admin", func(r chi.Router) {
		r.Post("/", admn.CreateAdmin(ctx))
	})

	r.logger.Info(fmt.Sprintf("init controller %s", op))
}

func (r *RestController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
