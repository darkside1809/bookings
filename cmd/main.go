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
	"github.com/darkside1809/bookings/internal/driver"
	"github.com/darkside1809/bookings/internal/handlers"
	"github.com/darkside1809/bookings/internal/helpers"
	"github.com/darkside1809/bookings/internal/models"
	"github.com/darkside1809/bookings/internal/render"
	"github.com/darkside1809/bookings/cmd/server"
)

// Main function start listen and serve a server at 0.0.0.0:9999
func main() {
	db, err := execute()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(server.App.MailChan)
	fmt.Println("Starting mail listener...")
	server.ListenForMail()

	host := "0.0.0.0"
	port := "9999"
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server.Routes(&server.App),
	}

	fmt.Printf("Server start listening at %s ;) Let's GOOO!\n", srv.Addr,)
	err = srv.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}

// execute function uses gob package for encoding and decoding data from models, 
// creates session and cookies, connects to PostgreSQL database,
// creates templates and set to them default pattern,
// uses NewRepo, NewHandlers, NewRenderer, NewHelpers to initialize the server to work properly
func execute() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	server.App.MailChan = mailChan

	server.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	server.App.InfoLog = server.InfoLog

	server.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	server.App.ErrorLog = server.ErrorLog

	server.Session = scs.New()
	server.Session.Lifetime = 24 * time.Hour
	server.Session.Cookie.Persist = true
	server.Session.Cookie.SameSite = http.SameSiteLaxMode
	server.Session.Cookie.Secure = server.App.InProduction

	server.App.Session = server.Session

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
	server.App.TemplateCache = tc
	server.App.UseCache = false

	repo := handlers.NewRepo(&server.App, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&server.App)
	helpers.NewHelpers(&server.App)

	return db, nil
}
