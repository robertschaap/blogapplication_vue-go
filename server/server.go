package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

type server struct {
	r *mux.Router
	port string
}

func (s *server) start() {
	s.routes()
	fmt.Println("Running on http://localhost" + s.port)
	log.Fatal(http.ListenAndServe(s.port, s.r))
}
