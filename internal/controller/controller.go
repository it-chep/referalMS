package controller

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"referalMS/internal/config"
	admin "referalMS/internal/controller/admin"
	"referalMS/internal/controller/dto"
	v1 "referalMS/internal/controller/v1"
	"referalMS/internal/domain/entity"
)

type UserService interface {
	RegisterNewUser(ctx context.Context)
}

type AdminService interface {
	GetWinners(ctx context.Context, dto dto.ExternalAdminDTO, filters dto.GetWinnersDTO) (winners []entity.Referal, err error)
}

type ReferalService interface {
	RegisterNewReferal(ctx context.Context, dto dto.ReferalUserDTO, adto dto.ExternalAdminDTO) (referalLink string, err error)
	GetReferalStatistic(ctx context.Context, dto dto.ReferalStatisticDTO, adto dto.ExternalAdminDTO) (allUsers, lastNDays int64, err error)
}

type RestController struct {
	router         *chi.Mux
	adminService   AdminService
	referalService ReferalService
	userService    UserService
	cfg            config.Config
	logger         *slog.Logger
}

func NewRestController(
	adminService AdminService,
	referalService ReferalService,
	userService UserService,
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

	admn := admin.NewAdmin(r.logger)

	apiV1Router.Post("/new_referal/", apiV1.CreateReferal(ctx))
	apiV1Router.Post("/new_user/", apiV1.CreateUser(ctx))
	apiV1Router.Post("/get_statistic/", apiV1.GetReferalStatistic(ctx))
	apiV1Router.Post("/get_winners/", apiV1.GetWinners(ctx))

	r.router.Route("/api", func(r chi.Router) {
		r.Mount("/v1", apiV1Router)
	})

	r.router.Route("/admin", func(r chi.Router) {
		r.Post("/", admn.CreateAdmin())
	})

	r.logger.Info(fmt.Sprintf("init controller %s", op))
}

func (r *RestController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
