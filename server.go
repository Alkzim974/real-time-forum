package main

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/handler"
	"text/template"

	"github.com/gorilla/mux"
)


func main() {
	database.InitDb()
    r := mux.NewRouter().StrictSlash(true)
    r.PathPrefix("/static/").Handler((http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))))
	r.HandleFunc("/", index).Methods("GET")
    r.HandleFunc("/home", handler.Home).Methods("GET")
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("POST")
	r.HandleFunc("/post", handler.Post).Methods("POST")

    fmt.Printf("Server Started on http://localhost%s/\n", ":8080")
    http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("static/index.html"))
	index.Execute(w, nil)
}
