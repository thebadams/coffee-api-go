package server

import "net/http"

func CreateServer() *http.Server {
	mux := createMux()
	s := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	return s
}
