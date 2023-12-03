package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/handlers"
	"github.com/manuel-valles/bookings-app.git/internal/models"
	"github.com/manuel-valles/bookings-app.git/internal/render"
)

const address = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// TODO: This should be based on environment variable instead
	app.InProduction = false

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Serving application on port %s\n", address)
	server := &http.Server{
		Addr:    address,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
