// Package handlers handle all api endpoint
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/shahariaazam/openapi-ninja/pkg/config"
)

// Status of the application
type Status struct {
	BuildID        string `json:"build_id"`
	CommitShaFull  string `json:"commit_sha_full"`
	CommitShaShort string `json:"commit_sha_short"`
	GitHubURL      string `json:"github_url"`
}

func getAppStatus(cfg config.Config) ([]byte, error) {
	s := Status{
		BuildID:        cfg.BuildID,
		CommitShaFull:  cfg.CommitSHAFull,
		CommitShaShort: cfg.CommitSHAShort,
		GitHubURL:      fmt.Sprintf("%s/commit/%s", cfg.GitHubRepositoryURL, cfg.CommitSHAFull),
	}

	return json.Marshal(s)
}

// StatusHandler handle request for /status endpoint
func StatusHandler(w http.ResponseWriter, r *http.Request, cfg config.Config) {
	w.Header().Set("Content-Type", "application/json")

	status, err := getAppStatus(cfg)
	if err != nil {
		logrus.WithError(err).Error("failed to get app status")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	if _, err := w.Write(status); err != nil {
		logrus.WithError(err).Error("failed to write app status response")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
}
