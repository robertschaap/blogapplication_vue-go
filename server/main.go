package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/comments/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/posts/{category}", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/posts/{id}", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/posts/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/users/new", helloWorldHandler).Methods("GET")
	r.HandleFunc("/api/users/login", helloWorldHandler).Methods("GET")

	port := ":3000"

	fmt.Println("Running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))

}

func main() {
	startServer()
}
