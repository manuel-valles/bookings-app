package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// Once all tests are done exit
	os.Exit(m.Run())
}

type mockHandler struct{}

func (mh *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
