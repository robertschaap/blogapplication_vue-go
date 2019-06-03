package controllers

import (
	"encoding/json"
	"net/http"

	"../api"
)

// GetUserLoginHandler does stuff
func GetUserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		ID				string	`json:"id"`
		FirstName	string	`json:"firstName"`
		LastName	string	`json:"lastName"`
		Token			string	`json:"authToken"`
	}

	user := User{
		ID: "0",
		FirstName: "robert",
		LastName: "schaap",
		Token: "mytoken",
	}

	res := api.Response{
		Status: "success",
		Data: user,
		Message: "",
	}

	res.JSON(w)
}


// CreateUserHandler does stuff
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
