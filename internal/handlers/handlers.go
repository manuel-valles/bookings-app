package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/models"
	"github.com/manuel-valles/bookings-app.git/internal/render"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	remoteIP := rp.App.Session.GetString(r.Context(), cookieIP)
	stringMap[cookieIP] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Standards(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "standards.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Suites(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "suites.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (rp *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := JSONResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
