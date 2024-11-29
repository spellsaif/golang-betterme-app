package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spellsaif/golang-betterme-app/internal/handlers"
	"github.com/spellsaif/golang-betterme-app/internal/middlewares"
)

func AuthRoutes(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", h.CreateUser)
	r.Post("/login", h.Login)
	r.With(middlewares.AuthMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to protected route, you will only see this if you are authenticated"))
	})

	return r
}
