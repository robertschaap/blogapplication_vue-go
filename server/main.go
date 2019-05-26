package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	c "./controllers"
)

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/comments/new", c.CreateCommentHandler).Methods("POST")
	r.HandleFunc("/api/posts", c.GetPostsHandler).Methods("GET")
	r.HandleFunc("/api/posts/{category}", c.GetPostsByCategoryHandler).Methods("GET")
	r.HandleFunc("/api/post/{id}", c.GetPostHandler).Methods("GET")
	r.HandleFunc("/api/post/new", c.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/users/new", c.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/login", c.GetUserLoginHandler).Methods("GET")

	port := ":3001"

	fmt.Println("Running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))

}

func main() {
	startServer()
}
