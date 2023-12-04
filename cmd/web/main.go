package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/driver"
	"github.com/manuel-valles/bookings-app.git/internal/handlers"
	"github.com/manuel-valles/bookings-app.git/internal/helpers"
	"github.com/manuel-valles/bookings-app.git/internal/models"
	"github.com/manuel-valles/bookings-app.git/internal/render"
)

const address = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Serving application on port %s\n", address)
	server := &http.Server{
		Addr:    address,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// TODO: This should be based on environment variable instead
	app.InProduction = false

	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.Reservation{})

	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=postgres")
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
