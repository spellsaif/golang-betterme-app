package routes

import (
	"github.com/go-chi/chi"
	"github.com/spellsaif/golang-betterme-app/internal/handlers"
)

func PostRoutes(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.GetPost)

	return r
}
