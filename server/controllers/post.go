package controllers

import (
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"net/http"

	"../api"
)

// GetPostHandler does stuff
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println(id)

	type User struct {
		FirstName	string	`json:"firstName"`
		LastName  string	`json:"lastName"`
		Bio				string	`json:"bio"`
	}

	type PostBody struct {
		Title	string	`json:"title"`
		Body  string	`json:"body"`
	}

	type Comment struct {
		FirstName	string	`json:"firstName"`
		LastName  string	`json:"lastName"`
		Body			string	`json:"body"`
	}

	type Post struct {
		Author		User			`json:"author"`
		Body			PostBody	`json:"body"`
		Comments	[]Comment	`json:"comments"`
	}

	author := User{ "robert", "schaap", "Robert is a true hero, he taught himself everything and has still remained humble. Very humble. More humble than you" }

	body := PostBody{
		Title: "That feeling when you're coding in a new framework",
		Body: "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
	}

	var comments []Comment
	comments = append(comments, Comment{ "robert", "schaap", "hi friend" })
	comments = append(comments, Comment{ "robert", "schaap", "hi friend" })
	comments = append(comments, Comment{ "robert", "schaap", "hi friend" })

	post := Post{ author, body, comments }

	res := api.Response{
		Status: "success",
		Data: post,
		Message: "",
	}

	res.JSON(w)
}


// CreatePostHandler does stuff
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}
