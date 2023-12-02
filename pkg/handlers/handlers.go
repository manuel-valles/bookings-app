package handlers

import (
	"net/http"

	"github.com/manuel-valles/bookings-app.git/pkg/config"
	"github.com/manuel-valles/bookings-app.git/pkg/models"
	"github.com/manuel-valles/bookings-app.git/pkg/render"
)

const cookieIP = "remote_ip"

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewHandlers(r *Repository) {
	Repo = r
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), cookieIP, remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	remoteIP := rp.App.Session.GetString(r.Context(), cookieIP)
	stringMap[cookieIP] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Standards(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "standards.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Suites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suites.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
