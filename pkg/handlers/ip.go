// Package handlers handle all api endpoint
package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// WhatIsMyIPHandler resolves IP address from http request
func WhatIsMyIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := RequestFromIP(r)

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

// RequestFromIP finds which IP made the request
func RequestFromIP(r *http.Request) string {
	var ip string

	ips := r.Header.Get("X-Forwarded-For")
	if ips != "" {
		ip = strings.Split(ips, ", ")[0]
	}

	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
