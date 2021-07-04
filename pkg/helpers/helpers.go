package helpers

import (
	// built in Golang packages
	"fmt"
	"net/http"
	"runtime/debug"
	// My own packages
	"github.com/darkside1809/bookings/pkg/config"
)

// app is a pointer/entry point to configuration of app
var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError log errors and set status for client side
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status ok", status)
	http.Error(w, http.StatusText(status), status)
}
// ServerError log errors and set status for server side
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
// IsAuthenticated checks if the given key is present in the session data
func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
