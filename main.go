// Package main is the main package
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/sirupsen/logrus"

	"github.com/shahariaazam/openapi-ninja/pkg/config"
	"github.com/shahariaazam/openapi-ninja/pkg/handlers"
	"github.com/shahariaazam/openapi-ninja/pkg/logging"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger := logging.NewLogger(cfg)

	logger.WithField("http_port", cfg.HTTPPort).Info("Listening on HTTP port")
	if err = http.ListenAndServe(":"+cfg.HTTPPort, httpRouters(cfg, logger)); err != nil {
		logger.WithField("http_port", cfg.HTTPPort).WithError(err).Error("failed to listen on HTTP port")
	}
}

func httpRouters(cfg config.Config, logger *logrus.Logger) *chi.Mux {
	httpRouter := chi.NewRouter()

	httpRouter.Use(middleware.RequestLogger(&logging.LogrusLogger{Logger: logger}))
	httpRouter.Use(httprate.LimitByIP(cfg.APIRequestLimitPerMinute, 1*time.Minute))
	httpRouter.Use(middleware.Heartbeat("/ping"))
	httpRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	httpRouter.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(w, r, cfg)
	})

	httpRouter.Get("/api/what-is-my-ip", handlers.WhatIsMyIPHandler)
	return httpRouter
}
