package handlers

import (
	"net/http"

	"github.com/darkside1809/bookings/pkg/config"
	"github.com/darkside1809/bookings/pkg/models"
	"github.com/darkside1809/bookings/pkg/render"
)

// Repo used by handlers
var Repo *Repository

// Repository holds App - pointer to a AppConfig structure
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository fot the handlers
func NewHandlers(r *Repository) {
	Repo = r 
}

// Home handler set handler to home route 
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About handler set handler to about route 
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}