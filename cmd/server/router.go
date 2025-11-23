package main

func (s *APIServer) registerRoutes() {
	// API List
	s.mux.HandleFunc("/apis/gadgets.muntashirislam.com/v1", APIRoot)

	// Gadgets Resource
	s.mux.HandleFunc("/apis/gadgets.muntashirislam.com/v1/gadgets", ListOrCreate)
}
