package controllers

import (
	"encoding/json"
	"net/http"
)

// GetPostHandler does stuff
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}


// CreatePostHandler does stuff
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
