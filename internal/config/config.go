package config

import (
	// built in Golang packages
	"log"
	"html/template"
	// External dependencies
	"github.com/alexedwards/scs/v2"
	// My own packages
	"github.com/darkside1809/bookings/internal/models"
)

// AppConfig holds the application's configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan		  chan models.MailData
}
