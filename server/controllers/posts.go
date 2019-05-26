package controllers

import (
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"net/http"
)

// GetPostsHandler doesnt do much yet
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	type Post struct {
		ID      int			`json:"id"`
		Title   string	`json:"title"`
		Author  string	`json:"author"`
	}

	var post []Post

	post = append(post, Post{ 0, "title", "author" })
	post = append(post, Post{ 1, "title", "author" })
	post = append(post, Post{ 2, "title", "author" })
	post = append(post, Post{ 3, "title", "author" })
	post = append(post, Post{ 4, "title", "author" })
	post = append(post, Post{ 5, "title", "author" })

	fmt.Println(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// GetPostsByCategoryHandler doesnt do much yet
func GetPostsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	fmt.Println(category)

	type Post struct {
		ID      int			`json:"id"`
		Title   string	`json:"title"`
		Author  string	`json:"author"`
	}

	var post []Post

	post = append(post, Post{ 0, "title", "author" })
	post = append(post, Post{ 1, "title", "author" })
	post = append(post, Post{ 2, "title", "author" })
	post = append(post, Post{ 3, "title", "author" })
	post = append(post, Post{ 4, "title", "author" })
	post = append(post, Post{ 5, "title", "author" })

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
