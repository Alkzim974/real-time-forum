package main

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/handler"
)

func main() {
	database.InitDb()
	// Association de la route "/" avec homeHandler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", handler.Register)

	fmt.Println("Serveur démarré sur le port 8080...")
	fmt.Println("http://localhost:8080")
	// Démarre l'écoute sur le port 8080
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page d’accueil !")
}
