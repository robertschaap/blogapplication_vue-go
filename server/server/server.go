package server

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	R *mux.Router
	Port string
}

// Start starts server
func (s *Server) Start() {
	s.routes()
	fmt.Println("Running on http://localhost" + s.Port)
	log.Fatal(http.ListenAndServe(s.Port, s.R))
}
