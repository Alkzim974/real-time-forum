package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJson(w http.ResponseWriter, status int, message any) {
	response, _ := json.Marshal(message)
	fmt.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
