package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/shahariaazam/openapi-ninja/pkg/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", handlers.PingHandler)

	r.Get("/api/what-is-my-ip", handlers.WhatIsMyIPHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
