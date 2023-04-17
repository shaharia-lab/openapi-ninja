package logging

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type AppEngineLogWriter struct{}

func (w *AppEngineLogWriter) Write(p []byte) (n int, err error) {
	ctx := appengine.NewContext(nil)
	log.Debugf(ctx, string(p))
	return len(p), nil
}
