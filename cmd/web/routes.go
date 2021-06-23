package main

import (
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/darkside1809/bookings/pkg/config"
	"github.com/darkside1809/bookings/pkg/handlers"
)

// Route set routes with its handler, multiplexer(middleware, sessionFuncs, secure requests), 
// and HTTP methods 
func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}