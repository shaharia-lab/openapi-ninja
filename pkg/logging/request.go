// Package logging provides logging functionality for the application
package logging

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is a wrapper around a logrus logger that implements the chi.Logger interface
type LogrusLogger struct {
	*logrus.Logger
}

// NewLogEntry implements the chi.Logger interface
func (l *LogrusLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := logrus.NewEntry(l.Logger)

	if rid := middleware.GetReqID(r.Context()); rid != "" {
		entry = entry.WithField("request_id", rid)
	}

	return &LogrusLogEntry{entry}
}

// LogrusLogEntry is a wrapper around a logrus Entry that implements the chi.LogEntry interface
type LogrusLogEntry struct {
	entry *logrus.Entry
}

// Write implements the chi.LogEntry interface
func (l *LogrusLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.entry.WithFields(logrus.Fields{
		"status":      status,
		"bytes":       bytes,
		"elapsed":     elapsed.Seconds(),
		"elapsed_str": elapsed.String(),
		"header":      header,
		"from_ip":     l.getRequestFromIPFromHeader(header),
	}).Info("request complete")
}

// Panic implements the chi.LogEntry interface
func (l *LogrusLogEntry) Panic(v interface{}, stack []byte) {
	l.entry.WithFields(logrus.Fields{
		"stack": string(stack),
	}).Errorf("%v", v)
}

func (l *LogrusLogEntry) getRequestFromIPFromHeader(h http.Header) string {
	ips := h.Get("X-Forwarded-For")
	if ips != "" {
		return strings.Split(ips, ", ")[0]
	}

	return ""
}
