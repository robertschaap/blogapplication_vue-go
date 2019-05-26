package main

import (
	"github.com/gorilla/mux"

	"./server"
)

func main() {
	s := server.Server{ R: mux.NewRouter(), Port: ":3001" }
	s.Start()
}
