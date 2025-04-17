package main

import (
	"app/pkg/config"
	"app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)      // CSRF protection
	mux.Use(SessionLoad) // Load and save session

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))
	return mux
}
