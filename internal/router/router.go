package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

type Router struct {
	router *chi.Mux
}

func NewRouter(logger *slog.Logger) *Router {
	const op = "router.router.NewRouter"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	return &Router{router: r}
}

func (r *Router) InitRouter() {
	api := r.router

	api.Route("/api_v1", func(r chi.Router) {
		r.Post("/new_referal", NewReferalHandler)
		r.Post("/new_user", NewUserHandler)
		r.Post("/get_statistic", GetStatisticHandler)
		r.Post("/get_winners", GetWinnersHandler)
	})

}

func NewReferalHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic for handling new referral
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic for handling new user
}

func GetStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic for getting statistics
}

func GetWinnersHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic for getting winners
}
