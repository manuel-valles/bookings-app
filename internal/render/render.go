package render

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/models"
)

var app *config.AppConfig
var functions = template.FuncMap{}
var pathPageTemplates = "./templates/*.page.tmpl"
var pathLayoutTemplates = "./templates/*.layout.tmpl"

func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplateData, r *http.Request) *models.TemplateData {
	data.Flash = app.Session.PopString(r.Context(), "flash")
	data.Warning = app.Session.PopString(r.Context(), "warning")
	data.Error = app.Session.PopString(r.Context(), "error")
	data.CSRFToken = nosurf.Token(r)
	return data
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) error {
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
		return errors.New("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)
	t.Execute(buf, data)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(pathPageTemplates)
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
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
