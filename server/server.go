package server

import "net/http"

func CreateServer() *http.Server {
	// mux := http.NewServeMux()

	// mux.Handle("GET /coffee", handler http.Handler)
	s := &http.Server{
		Addr: ":3000",
	}
	return s
}
