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
	"github.com/jung-kurt/gofpdf"
	"github.com/justinas/nosurf"
	// My own packages
	"github.com/darkside1809/bookings/internal/config"
	"github.com/darkside1809/bookings/internal/models"
)

// Functions variable holds whole functionality of our template (html data)
var functions = template.FuncMap{
	"humanDate":  HumanDate,
	"formatDate": FormatDate,
	"iterate":    Iterate,
	"add":        Add,
}
// app is a pointer/entry point to configuration of app
var app *config.AppConfig
var pathToTemplates = "./templates"

func Add(a int, b int) int {
	return a + b 
}

// Iterate returns a slice of ints starting at 1, going to count
func Iterate(count int) []int {
	var i int
	var items []int
	for i = 0; i < count; i++ {
		items = append(items, i)
	}
	return items
}

// NewRenderer create new template for handlers
func NewRenderer(a *config.AppConfig) {
	app = a
}

// HumanDate returns formated time YYYY-MM-DD
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDate formats given date to a string
func FormatDate(t time.Time, f string) string {
	return t.Format(f)
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

// NewReport returns formated pdf
func NewReport(head string) *gofpdf.Fpdf{
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, head)
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)

	return pdf
}

// Header set header to pdf file
func Header(pdf *gofpdf.Fpdf, headerText []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)

	for _, str := range headerText {
		pdf.CellFormat(40, 10, str, "1", 0, "", true, 0, "")
	}
	pdf.Ln(-1)

	return pdf
}

// Table creates table with its own style
func Table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)

	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		for i, str := range line {
			pdf.CellFormat(40, 10, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
	pdf.Ln(-1)

	return pdf
}
