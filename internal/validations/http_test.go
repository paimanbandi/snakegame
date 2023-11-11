package validations

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsGetMethod(t *testing.T) {
	// Test for GET method
	req, err := http.NewRequest("GET", "/new", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IsGetMethod)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for GET: got %v want %v",
			status, http.StatusOK)
	}

	// Test for non-GET method (e.g., POST)
	req, err = http.NewRequest("POST", "/new", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for non-GET: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestIsPostMethod(t *testing.T) {
	// Test for POST method
	req, err := http.NewRequest("POST", "/validate", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IsPostMethod)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for POST: got %v want %v",
			status, http.StatusOK)
	}

	// Test for non-POST method (e.g., GET)
	req, err = http.NewRequest("GET", "/validate", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for non-POST: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
