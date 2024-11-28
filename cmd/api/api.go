package main

import (
	"net/http"

	"github.com/spellsaif/golang-betterme-app/internal/routes"
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

	r := routes.NewRouter()
	http.ListenAndServe(a.addr, r)
}
