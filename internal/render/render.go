package render

import (
	"bytes"
	//"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/models"
	"github.com/justinas/nosurf"
)

// Functions variable holds whole functionality of our template (html data)
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates create new template for handlers
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData set and return default struct of data 
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders data to client
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template 
	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}
	
	t, ok := tc[tmpl]
	if !ok {
		log.Print("Couldn't get template, there is no one of them")
		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Print(err)
		return
	}
}

// CreateTemplateCache find out all exact extensions and set to them default pattern structure
func CreateTemplateCache() (map[string]*template.Template, error){
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}