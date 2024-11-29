package main

import (
	"log"
	"net/http"

	"github.com/spellsaif/golang-betterme-app/internal/routes"
	"github.com/spellsaif/golang-betterme-app/internal/storage"
)

type Api struct {
	addr string
}

func NewApi(addr string) *Api {

	return &Api{
		addr: addr,
	}
}

func (a *Api) Run() {

	db, err := storage.New()

	if err != nil {
		log.Fatal("failed to connect to database")
	}

	defer db.Db.Close()

	r := routes.NewRouter(db)
	http.ListenAndServe(a.addr, r)
}
