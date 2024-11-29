package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spellsaif/golang-betterme-app/internal/handlers"
	"github.com/spellsaif/golang-betterme-app/internal/middlewares"
	"github.com/spellsaif/golang-betterme-app/internal/storage"
)

func NewRouter(db *storage.Sqlite) *chi.Mux {

	//hanlders
	h := handlers.NewHandler(db)

	//creates new MUX(Router) for handling our endpoints
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	//Using logger middleware to log useful information
	r.Use(middleware.Logger)

	//Routes

	//health check route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working..."))
	})

	r.With(middlewares.AuthMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to protected route, you will only see this if you are authenticated"))
	})

	//subroute
	authRoute := chi.NewRouter()
	authRoute.Post("/create", h.CreateUser)
	authRoute.Post("/login", h.Login)

	//now mounting it to main route
	r.Mount("/auth", authRoute)

	//not found round
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("<h1>Sorry you landed to wrong planet!<h1>"))
	})

	return r
}
