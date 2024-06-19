package main

import (
	"log"

	"github.com/thebadams/coffee-api-go/server"
)

func main() {
	s := server.CreateServer()
	log.Printf("Server Starting on PORT %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
