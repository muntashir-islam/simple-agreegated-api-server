package main

import (
	"log"
)

func main() {
	server := New()
	log.Println("Starting Aggregated API Server on :8443")
	server.RunTLS(":8443")
}
