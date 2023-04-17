// Package logging provides logging functionality for the application
package logging

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// AppEngineLogWriter provides log writer for App Engine
type AppEngineLogWriter struct{}

// Write the log to app engine logger
func (w *AppEngineLogWriter) Write(p []byte) (n int, err error) {
	ctx := appengine.NewContext(nil)
	log.Debugf(ctx, string(p))
	return len(p), nil
}
