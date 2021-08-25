package main

import (
	"github.com/go-chi/chi/v5"
	"go_blog/internal/config"
	"go_blog/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)

	return mux
}