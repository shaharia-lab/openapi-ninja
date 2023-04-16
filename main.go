package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", pingHandler)

	r.Get("/api/what-is-my-ip", myIPHandler)

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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func myIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}

	format := r.URL.Query().Get("format")
	switch format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"ip": "%s"}`, ip)
	case "text":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", ip)
	case "xml":
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, "<ip>%s</ip>", ip)
	default:
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"ip": "%s"}`, ip)
	}
}
