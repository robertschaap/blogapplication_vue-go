package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"./controllers"
)

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/comments/new", controllers.CreateCommentHandler).Methods("POST")
	r.HandleFunc("/api/posts", controllers.GetPostsHandler).Methods("GET")
	r.HandleFunc("/api/posts/{category}", controllers.GetPostsByCategoryHandler).Methods("GET")
	r.HandleFunc("/api/post/{id}", controllers.GetPostHandler).Methods("GET")
	r.HandleFunc("/api/post/new", controllers.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/users/new", controllers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/login", controllers.GetUserLoginHandler).Methods("GET")

	port := ":3001"

	fmt.Println("Running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))

}

func main() {
	startServer()
}
