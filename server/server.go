package server

import "net/http"

func CreateServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", hellohandler)
	s := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	return s
}
