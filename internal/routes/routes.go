package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spellsaif/golang-betterme-app/internal/handlers"
	"github.com/spellsaif/golang-betterme-app/internal/middlewares"
)

func NewRouter() *chi.Mux {
	//creates new MUX(Router) for handling our endpoints
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	//Using logger middleware to log useful information
	r.Use(middleware.Logger)

	//Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working..."))
	})

	//not found round
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("<h1>Sorry you landed to wrong planet!<h1>"))
	})

	//subroute
	authRoute := chi.NewRouter()

	//subroute routes
	authRoute.With(middlewares.AuthMiddleware).Get("/", handlers.GetUser)

	authRoute.Post("/", handlers.CreateUser)

	//now mounting it to main route
	r.Mount("/auth", authRoute)

	return r
}
