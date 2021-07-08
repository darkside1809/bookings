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
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/color"
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

// BuildHeading creates header a pdf file
func BuildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("static/images/about.jpg", props.Rect{
					Center: true,
					Percent: 90,
				})
				if err != nil {
					log.Print(err)
				}
			})
		})
	})
}

// BuildList creates a table
func BuildList(m pdf.Maroto, contents [][]string) {
	tableHeadings := []string{"Name", "Room", "Arrival", "Departure", "Email", "Phone"}
	
	lightPurpleColor := getLightPurpleColor()
	m.SetBackgroundColor(getDarkPurpleColor())

	m.Row(12, func() {
		m.Col(12, func() {
			m.Text("Reservation Summary", props.Text{
				Top: 		2,
				Size: 	15,
				Color: 	color.NewWhite(),
				Family: 	consts.Courier,
				Style: 	consts.Bold,
				Align: 	consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size: 11,
			GridSizes: []uint{2, 2, 2, 2, 3, 3},
		},
		ContentProp: props.TableListContent{
			Size: 9,
			GridSizes: []uint{2, 2, 2, 2, 3, 3},
		},
		Align: consts.Left,
		HeaderContentSpace: 1,
		Line: false,
		AlternatedBackground: &lightPurpleColor,
	})
}

// getTealColor get definite color 
func getTealColor() color.Color {
	return color.Color{
		Red: 3,
		Green: 166,
		Blue: 166,
	}
}

// getLightPurpleColor get definite color 
func getLightPurpleColor() color.Color {
	return color.Color{
		Red: 210,
		Green: 200,
		Blue: 230,
	}
}

// getDarkPurpleColor get definite color 
func getDarkPurpleColor() color.Color {
	return color.Color{
		Red: 88,
		Green: 88,
		Blue: 99,
	}
}
