package server

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

	route := r.URL
	log.Printf("%s HIT", route)
}
