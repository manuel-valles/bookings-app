package render

import (
	"net/http"
	"testing"

	"github.com/manuel-valles/bookings-app.git/internal/models"
)

func TestNewRenderer(t *testing.T) {
	NewRenderer(app)
}

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Fatal(err)
	}

	const mockedFlashValue = "testing flash"

	// Check context get updated
	session.Put(r.Context(), "flash", mockedFlashValue)

	result := AddDefaultData(&td, r)
	if result.Flash != mockedFlashValue {
		t.Errorf(`Flash value of "%s" not found in session`, mockedFlashValue)
	}

}

func TestTemplate(t *testing.T) {
	pathPageTemplates = "../../templates/*.page.tmpl"
	pathLayoutTemplates = "../../templates/*.layout.tmpl"

	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	// Mock request
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var mw mockedWriter

	err = Template(&mw, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser", err)
	}

	err = Template(&mw, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Rendered template that does not exist")
	}

}

func TestCreateTemplateCache(t *testing.T) {
	pathPageTemplates = "../../templates/*.page.tmpl"
	pathLayoutTemplates = "../../templates/*.layout.tmpl"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
