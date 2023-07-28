package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Specify who is allowed to connect with Chi/Cors
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		// TODO: Look at what Allow Credentials does
		AllowCredentials: true,
		MaxAge: 300,
	}))

	// To check if this service is alive
	mux.Use(middleware.Heartbeat("/ping"))

	// Adding Route Handlers

	mux.Post("/", app.Broker)

	mux.Post("/handle", app.HandleSubmission)

	return mux
}