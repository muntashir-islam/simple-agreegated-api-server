package main

import (
	"net/http"
)

type APIServer struct {
	mux *http.ServeMux
}

func New() *APIServer {
	s := &APIServer{
		mux: http.NewServeMux(),
	}

	s.registerRoutes()
	return s
}

func (s *APIServer) RunTLS(addr string) {
	cert := "/certs/tls.crt"
	key := "/certs/tls.key"
	http.ListenAndServeTLS(addr, cert, key, s.mux)
}
