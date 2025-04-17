package main

import (
	"net/http"

	"github.com/PedroGabrielBHZ/bookings/pkg/config"
	"github.com/PedroGabrielBHZ/bookings/pkg/handlers"

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

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
