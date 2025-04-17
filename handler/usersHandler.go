package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"sort"
)

func displayUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetAllUsers()
	
	if len(users) == 0 {
		fmt.Fprintln(w, "Aucun utilisateur trouvé.")
	} else {
		// Tri des utilisateurs par ordre alphabétique (selon le Nickname)
		sort.Slice(users, func(i, j int) bool {
			return users[i].Nickname < users[j].Nickname
		})
		
		fmt.Fprintln(w, "Utilisateurs:")
		for _, user := range users {
			fmt.Fprintln(w, " ", user.Nickname)
		}
	}
}