package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// HelloResponse - test struct
type HelloResponse struct {
	Message string `json:"message"`
}

// RandomNumberResponse -
type RandomNumberResponse struct {
	Message int `json:"message"`
}

// HelloHandler - test handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello",
	}
	json.NewEncoder(w).Encode(response)
	return
}

// RandomNumberHandler - return message with random int within the range 1-100,000
func RandomNumberHandler(w http.ResponseWriter, r *http.Request) {
	response := RandomNumberResponse{
		Message: rand.Intn(100000),
	}
	json.NewEncoder(w).Encode(response)
	return
}
