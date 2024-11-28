package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spellsaif/golang-betterme-app/internal/handlers"
)

func NewRouter() *chi.Mux {
	//creates new MUX(Router) for handling our endpoints
	r := chi.NewRouter()

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
	authRoute.Get("/", handlers.GetUser)

	authRoute.Post("/", handlers.CreateUser)

	//now mounting it to main route
	r.Mount("/auth", authRoute)

	return r
}
