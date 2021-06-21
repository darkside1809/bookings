package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/darkside1809/bookings/pkg/config"
	"github.com/darkside1809/bookings/pkg/handlers"
	"github.com/darkside1809/bookings/pkg/render"
)

// App holds AppConfig structure
var app config.AppConfig
// Session holds pointer to SessionManager structure from external package
var session *scs.SessionManager

// Main function starts listening a server at host, port
func main() {
	// Change this to true when app is in production
	// But in the development mode we set it to false 
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction 

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Print(err)
		return
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	host := "0.0.0.0"
	port := "9999"
	
	srv := &http.Server {
		Addr: net.JoinHostPort(host, port),
		Handler: routes(&app),
	}

	fmt.Printf("Server start listening at %s\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}

