package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("not_field_yet")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData.Add("first_key", "first_value")

	form = New(postedData)

	has = form.Has("first_key")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid form when it is valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	form := New(r.PostForm)

	form.Required("first_key", "second_key")
	if form.Valid() {
		t.Error("go valid form when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("first_key", "first_value")
	postedData.Add("second_key", "second_value")

	r, _ = http.NewRequest("POST", "/any", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("first_key")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("not_field_yet", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("not_field_yet")
	if isError == "" {
		t.Error("should have error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("first_key", "first_value")

	form = New(postedData)

	minLength := 100
	form.MinLength("first_key", minLength)
	if form.Valid() {
		t.Errorf("shows min length of %d met when data is shorter", minLength)
	}

	postedData = url.Values{}
	postedData.Add("second_key", "second_value")
	form = New(postedData)

	minLength = 1
	form.MinLength("second_key", minLength)
	if !form.Valid() {
		t.Errorf("shows min length of %d is not met when it is", minLength)
	}

	isError = form.Errors.Get("second_key")
	if isError != "" {
		t.Error("should not have error but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("not_field_yet")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "a@a.a")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when it is valid")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "no_valid_email")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
