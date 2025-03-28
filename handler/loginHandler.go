package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Identifiant string `json:"identifiant"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page de connexion !")
	userLogin := LoginResponse{}
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email, err := mail.ParseAddress(regResp.Email)
	
}
