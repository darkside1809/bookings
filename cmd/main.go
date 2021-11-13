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
	"github.com/joho/godotenv"

	// My own packages
	"github.com/darkside1809/bookings/cmd/server"
	"github.com/darkside1809/bookings/internal/driver"
	"github.com/darkside1809/bookings/internal/handlers"
	"github.com/darkside1809/bookings/internal/helpers"
	"github.com/darkside1809/bookings/internal/models"
	"github.com/darkside1809/bookings/internal/render"
)

func main() {

	db, err := execute()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(server.App.MailChan)
	fmt.Println("Starting mail listener...")
	server.ListenForMail()

	host, port := os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server.Routes(&server.App),
	}

	fmt.Printf("Server start listening at %s ;) Let's Gooo!\n", srv.Addr,)
	if err = srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}

// execute function uses gob package for encoding and decoding data from models, 
// creates session and cookies, connects to PostgreSQL database,
// creates templates and set to them default pattern,
// uses NewRepo, NewHandlers, NewRenderer, NewHelpers to initialize the server to work properly
func execute() (*driver.DB, error) {
	
	setServerConfigs()
	
	// Connect to database
	log.Println("Connecting to database...")

	dsn := setDBconfigs()
	db, err := driver.ConnectSQL(dsn)
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

func setDBconfigs() (dsn string){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't load env vars")
		return ""
	}
	
	dbConfigs := models.DBconfigs {
		Name: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
	}

	dsn = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbConfigs.Host, dbConfigs.Port, dbConfigs.Name, dbConfigs.User, dbConfigs.Password)

	return dsn
}

func setServerConfigs() {
	
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
}
