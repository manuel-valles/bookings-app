package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var handlersTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"standards", "/standards", "GET", []postData{}, http.StatusOK},
	{"suites", "/suites", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-02"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Manu"},
		{key: "last_name", value: "Kem"},
		{key: "email", value: "a@a.a"},
		{key: "phone", value: "123-456-7890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	mockedServer := httptest.NewTLSServer(routes)
	defer mockedServer.Close()

	for _, currentTest := range handlersTests {
		if currentTest.method == "GET" {
			resp, err := mockedServer.Client().Get(mockedServer.URL + currentTest.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != currentTest.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", currentTest.name, currentTest.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}

			for _, data := range currentTest.params {
				values.Add(data.key, data.value)
			}

			resp, err := mockedServer.Client().PostForm(mockedServer.URL+currentTest.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != currentTest.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", currentTest.name, currentTest.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
