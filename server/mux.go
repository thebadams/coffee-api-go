package server

import "net/http"

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", hellohandler)
	mux.HandleFunc("POST /coffee", newCoffeeHandler)
	mux.HandleFunc("GET /coffee", getAllCoffeeHandler)
	return mux
}
