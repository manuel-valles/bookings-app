package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/manuel-valles/bookings-app.git/pkg/config"
	"github.com/manuel-valles/bookings-app.git/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData, r *http.Request) *models.TemplateData {
	data.CSRFToken = nosurf.Token(r)
	return data
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)
	t.Execute(buf, data)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	pathPageTemplates := "./templates/*.page.tmpl"
	pathLayoutTemplates := "./templates/*.layout.tmpl"
	cache := map[string]*template.Template{}

	// Get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(pathPageTemplates)
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob(pathLayoutTemplates)
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			t, err = t.ParseGlob(pathLayoutTemplates)
			if err != nil {
				return cache, err
			}
		}
		cache[name] = t
	}

	return cache, nil
}
