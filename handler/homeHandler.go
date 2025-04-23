package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/utils"
)


func Home(w http.ResponseWriter, r *http.Request, hub *utils.Hub) {
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
	
	userlist := hub.GetOnlineUsers()
		if len(userlist) == 0 {
		fmt.Fprintln(w, "Aucun utilisateur trouvé.")
	} else {
		fmt.Fprintln(w, "Voici les utilisateurs :")
		for _, user := range userlist {
			fmt.Fprintf(w, "Pseudo: %s\n", user)
		}
	}

}
