package render

import (
	// built in Golang packages
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	// External packages/dependencies
	"github.com/justinas/nosurf"
	// My own packages
	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/models"
)

// Functions variable holds whole functionality of our template (html data)
var functions = template.FuncMap{
	"humanDate": HumanDate,
}
// app is a pointer/entry point to configuration of app
var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer create new template for handlers
func NewRenderer(a *config.AppConfig) {
	app = a
}
// HumanDate returns formated time YYYY-MM-DD
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// AddDefaultData set and return default struct of data
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}

	return td
}

// Template renders data to client
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		return errors.New("Can't take template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// CreateTemplateCache find out all exact extensions and set to them default pattern structure
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
