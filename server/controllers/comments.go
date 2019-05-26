package controllers

import (
	"encoding/json"
	"net/http"
)

// CreateCommentHandler does stuff
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
