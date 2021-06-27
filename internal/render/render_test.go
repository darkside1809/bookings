package render

import (
	"testing"

	"github.com/darkside1809/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww writer

	err = RenderTemplate(&ww, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	err = RenderTemplate(&ww, r, "non-existing.html", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that doesn't exist")
	}
}

func TestNewTemplate(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}