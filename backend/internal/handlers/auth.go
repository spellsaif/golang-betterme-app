package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spellsaif/golang-betterme-app/internal/models"
	"github.com/spellsaif/golang-betterme-app/internal/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)

	}

	id, err := h.Db.CreateUser(&user)

	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	fmt.Printf("type of user: %T\n", user)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(map[string]interface{}{"id": id})

	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	json.NewDecoder(r.Body).Decode(&user)

	result, err := h.Db.FindUserByUsername(user.Username)

	if err != nil {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)

	}

	if result.Password != user.Password {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)
	}

	tokenString, err := utils.CreateToken(user.Username)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
