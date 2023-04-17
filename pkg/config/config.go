// Package config load and process application configuration
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config store configuration for application
type Config struct {
	HTTPPort            string `envconfig:"PORT" default:"8080"`
	AppVersion          string `envconfig:"GAE_VERSION"`
	BuildID             string `envconfig:"BUILD_ID"`
	CommitSHAFull       string `envconfig:"COMMIT_SHA"`
	CommitSHAShort      string `envconfig:"SHORT_SHA"`
	GoogleCloudProject  string `envconfig:"GOOGLE_CLOUD_PROJECT"`
	RunningOnAppEngine  bool
	GitHubRepositoryURL string `envconfig:"GITHUB_REPOSITORY" default:"https://github.com/shahariaazam/openapi-ninja"`

	APIRequestLimitPerMinute int `envconfig:"API_REQUEST_LIMIT_PER_MINUTE" default:"10"`
}

// Load configuration
func Load() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("error loading environment variables")
	}

	config.RunningOnAppEngine = isAppEngine()

	return config, nil
}

func isAppEngine() bool {
	return false
}
