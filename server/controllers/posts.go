package controllers

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"

	"../api"
)

// GetPostsHandler doesnt do much yet
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	type Post struct {
		ID      int			`json:"id"`
		Title   string	`json:"title"`
		Author  string	`json:"author"`
	}

	var posts []Post

	posts = append(posts, Post{ 0, "title", "author" })
	posts = append(posts, Post{ 1, "title", "author" })
	posts = append(posts, Post{ 2, "title", "author" })
	posts = append(posts, Post{ 3, "title", "author" })
	posts = append(posts, Post{ 4, "title", "author" })
	posts = append(posts, Post{ 5, "title", "author" })

	res := api.Response{
		Status: "success",
		Data: posts,
		Message: "",
	}

	res.JSON(w)
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

	var posts []Post

	posts = append(posts, Post{ 0, "title", "author" })
	posts = append(posts, Post{ 1, "title", "author" })
	posts = append(posts, Post{ 2, "title", "author" })
	posts = append(posts, Post{ 3, "title", "author" })
	posts = append(posts, Post{ 4, "title", "author" })
	posts = append(posts, Post{ 5, "title", "author" })

	res := api.Response{
		Status: "success",
		Data: posts,
		Message: "",
	}

	res.JSON(w)
}
