package server

import (
	// built in Golang packages
	"net/http"
	"log"
	// External packages/dependencies
	"github.com/justinas/nosurf"
	"github.com/alexedwards/scs/v2"
	// My own packages
	"github.com/darkside1809/bookings/internal/helpers"
	"github.com/darkside1809/bookings/internal/config"
)

// App holds AppConfig structure
var App config.AppConfig
// Info and errors for proper error handling
var InfoLog *log.Logger
var ErrorLog *log.Logger
// Session holds pointer to SessionManager structure from external package
var Session *scs.SessionManager

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return Session.LoadAndSave(next)
}

// Auth checks user's authentication,
// -if IsAuthenticated returns true, user will be logged in
// -if IsAuthenticated returns false, user won't be able to log in
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			Session.Put(r.Context(), "error", "Login first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
