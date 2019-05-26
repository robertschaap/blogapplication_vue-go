package controllers

import (
	"encoding/json"
	"net/http"
)

// GetUserLoginHandler does stuff
func GetUserLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}


// CreateUserHandler does stuff
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
