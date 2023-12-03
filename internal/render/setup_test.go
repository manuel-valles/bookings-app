package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/manuel-valles/bookings-app.git/internal/config"
	"github.com/manuel-valles/bookings-app.git/internal/models"
)

var session *scs.SessionManager
var mockedApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	// Test environment
	mockedApp.InProduction = false

	mockedApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	mockedApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = mockedApp.InProduction

	mockedApp.Session = session

	app = &mockedApp

	os.Exit(m.Run())
}

type mockedWriter struct{}

func (tw *mockedWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *mockedWriter) WriteHeader(i int) {}

func (tw *mockedWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
