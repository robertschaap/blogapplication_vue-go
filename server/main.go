package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"./controllers"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/comments/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/posts", controllers.GetPostsHandler).Methods("GET")
	r.HandleFunc("/api/posts/{category}", controllers.GetPostsByCategoryHandler).Methods("GET")
	r.HandleFunc("/api/post/{id}", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/post/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/users/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/users/login", helloWorldHandler).Methods("GET")

	port := ":3001"

	fmt.Println("Running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))

}

func main() {
	startServer()
}
