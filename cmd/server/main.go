package main

import (
	"log"
)

func main() {
	server := apiserver.New()
	log.Println("Starting Aggregated API Server on :8443")
	server.RunTLS(":8443")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8443", nil))
}
