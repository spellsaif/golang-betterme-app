package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spellsaif/golang-betterme-app/internal/models"
	"github.com/spellsaif/golang-betterme-app/internal/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("User created"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("The value of user %v\n", user)

	if user.Username == "saif" || user.Password == "ali" {
		tokenString, err := utils.CreateToken(user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No user found")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Credentials")
	}

}
