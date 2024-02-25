package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Router struct {
	router *chi.Mux
}

func NewRouter() *Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	return &Router{router: r}
}

func (r *Router) InitRouter() {
	api := r.router.Group("/api")
	api.Post("/new_referal", NewReferalHandler)
	api.Post("/new_user", NewUserHandler)
	api.Post("/get_statistic", GetStatisticHandler)
	api.Post("/get_winners", GetWinnersHandler)
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
