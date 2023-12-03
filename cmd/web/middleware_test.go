package main

import (
	"net/http"
	"testing"
)

func TestNoServe(t *testing.T) {
	var mh mockHandler

	handler := NoSurf(&mh)

	switch result := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", result)
	}
}

func TestSessionLoad(t *testing.T) {
	var mh mockHandler

	handler := NoSurf(&mh)

	switch result := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", result)
	}
}
