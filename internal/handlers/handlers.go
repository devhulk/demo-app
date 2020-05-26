package handlers

import (
	"encoding/json"
	"net/http"
)

// HelloResponse - test struct
type HelloResponse struct {
	Message string `json:"message"`
}

// HelloHandler - test handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello",
	}
	json.NewEncoder(w).Encode(response)
	return
}
