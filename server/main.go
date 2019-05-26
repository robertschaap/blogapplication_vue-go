package main

import (
	"github.com/gorilla/mux"
)

func main() {
	s := server{ r: mux.NewRouter(), port: ":3001" }
	s.start()
}
