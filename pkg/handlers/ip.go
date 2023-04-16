package handlers

import (
	"fmt"
	"net/http"
)

func WhatIsMyIPHandler(w http.ResponseWriter, r *http.Request) {
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
