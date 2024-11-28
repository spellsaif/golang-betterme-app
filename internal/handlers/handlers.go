package handlers

import (
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("User created"))
}
