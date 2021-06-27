package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/handlers"
	"github.com/darkside1809/bookings/internal/helpers"
	"github.com/darkside1809/bookings/internal/models"
	"github.com/darkside1809/bookings/internal/render"
)

// App holds AppConfig structure
var app config.AppConfig

// Info and errors handling
var infoLog  *log.Logger
var errorLog *log.Logger
// Session holds pointer to SessionManager structure from external package
var session *scs.SessionManager

// Main function starts listening a server at host, port
func main() {
	err := execute()
	if err != nil {
		log.Fatal(err)
	}
	host := "0.0.0.0"
	port := "9999"

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: routes(&app),
	}

	fmt.Printf("Server start listening at %s\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}

func execute() error {
	gob.Register(models.Reservation{})

	// Change this to true when app is in production
	// But in the development mode we set it to false
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
