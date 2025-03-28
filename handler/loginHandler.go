package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"real-time-forum/database"
	"real-time-forum/variables"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Identifiant string `json:"identifiant"`
	Password    string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page de connexion !")
	var user *variables.User
	userLogin := LoginResponse{}
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email, err := mail.ParseAddress(userLogin.Identifiant)
	if err != nil {
		user = database.GetUserByNickname(userLogin.Identifiant)
	} else {
		user = database.GetUserByEmail(email.Address)
	}
	if user.Email == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userLogin.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	} else {
		fmt.Fprintln(w, "Login successful")
	}
}
