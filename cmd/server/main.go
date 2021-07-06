package main

import (
	// built in Golang packages
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	// External packages/dependencies
	"github.com/alexedwards/scs/v2"
	// My own packages
	"github.com/darkside1809/bookings/pkg/config"
	"github.com/darkside1809/bookings/pkg/driver"
	"github.com/darkside1809/bookings/pkg/handlers"
	"github.com/darkside1809/bookings/pkg/helpers"
	"github.com/darkside1809/bookings/pkg/models"
	"github.com/darkside1809/bookings/pkg/render"
)

// App holds AppConfig structure
var app config.AppConfig

// Info and errors handling
var infoLog *log.Logger
var errorLog *log.Logger

// Session holds pointer to SessionManager structure from external package
var session *scs.SessionManager

// Main function starts listening a server at host, port
func main() {
	db, err := execute()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)
	fmt.Println("Starting mail listener...")
	listenForMail()

	host := "0.0.0.0"
	port := "9999"
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: routes(&app),
	}

	fmt.Printf("Server start listening at %s ;) Let's GOOO!\n", srv.Addr,)
	err = srv.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}

func execute() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

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

	// Connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=nekruz1809")
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache")
		return nil, err
	}
	log.Println("Creating templates...")
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
