package main

import (
	"testing"

	"github.com/darkside1809/bookings/pkg/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:

	default:
		t.Errorf("type is not *chi.Mux, but %T", v)
	}
}
