package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/variables"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterResponse struct {
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page d’inscription !")
	userInfo := RegisterResponse{}
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := uuid.NewV4()
	fmt.Println(userInfo)

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)

	user := &variables.User{
		ID:        uuid.String(),
		Nickname:  userInfo.Nickname,
		Age:       userInfo.Age,
		Gender:    userInfo.Gender,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Password:  bcryptPassword,
	}

	database.InsertUser(user)
}
