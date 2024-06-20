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

func newCoffeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Coffee Added To Database")
	route := r.URL
	m := r.Method
	log.Printf("%s HIT with %s, New Coffee Created", route, m)
}

func getAllCoffeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Coffees Returned")
	route := r.URL
	method := r.Method
	log.Printf("%s HIT with %s", route, method)
}
