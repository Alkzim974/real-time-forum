package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
)


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page d’accueil !")
	posts :=database.GetpostHome()
	
	if len(posts) == 0 {
		fmt.Fprintln(w, "Aucun post trouvé.")
	} else {
		fmt.Fprintln(w, "Voici les posts :")
		for _, post := range posts {
			fmt.Fprintf(w, "Category: %s, Title: %s, Content: %s\n", post.Category, post.Title, post.Content)
		}
	}
	displayUsers(w, r)
}
