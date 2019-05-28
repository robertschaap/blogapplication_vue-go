package api

import (
	"encoding/json"
	"net/http"
)

// Response struct provides a standardised wrapper for api responses
type Response struct {
	Status	 string 		  `json:"status"`
	Data		 interface{}  `json:"data"`
	Message	 string 			`json:"message"`
}

// JSON method sets json content type and encodes the response
func (r *Response) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
