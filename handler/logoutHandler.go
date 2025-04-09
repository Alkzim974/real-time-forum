package handler

import (
	"net/http"
)


func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	cookie.MaxAge = -1 // Supprime le cookie
	cookie.Value = ""
	http.SetCookie(w, cookie)
}

