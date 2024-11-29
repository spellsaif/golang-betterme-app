package handlers

import (
	"github.com/spellsaif/golang-betterme-app/internal/storage"
)

type Handler struct {
	Db *storage.Sqlite
}

func NewHandler(db *storage.Sqlite) *Handler {
	return &Handler{
		Db: db,
	}
}
