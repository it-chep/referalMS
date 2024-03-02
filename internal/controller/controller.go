package controller

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	admin "referalMS/internal/controller/admin"
	v1 "referalMS/internal/controller/v1"
)

type RestController struct {
	router *chi.Mux
	logger *slog.Logger
}

func NewRestController(logger *slog.Logger) *RestController {
	const op = "controller.controller.NewRouter"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	logger.Info(fmt.Sprintf("init router %s", op))

	return &RestController{router: r, logger: logger}
}

func (r *RestController) InitController(ctx context.Context) {
	const op = "controller.controller.InitController"
	apiV1Router := chi.NewRouter()
	apiV1 := v1.NewApiV1(r.logger)
	apiV1Router.Use(apiV1.GetAdminMiddleware)

	admn := admin.NewAdmin(r.logger)

	apiV1Router.Post("/new_referal/", apiV1.CreateReferal())
	apiV1Router.Post("/new_user/", apiV1.CreateUser())
	apiV1Router.Post("/get_statistic/", apiV1.GetReferalStatistic())
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
