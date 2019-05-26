package api

import (
	"encoding/json"
	"net/http"
)

// Response is an Response
type Response struct {
	Status	 string 		  `json:"status"`
	Data		 interface{}  `json:"data"`
	Message	 string 			`json:"message"`
}

// JSON is a Json
func (r *Response) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
