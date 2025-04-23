package handler

import (
	"net/http"
	"real-time-forum/database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	posts := database.GetpostHome()

	if len(posts) == 0 {
		RespondJson(w, http.StatusNotFound, map[string]any{
			"error": "No Posts Found",
		})
	} else {
		RespondJson(w, http.StatusOK, map[string]any{
			"Posts": posts,
		})
	}
	//displayUsers(w, r)
}
