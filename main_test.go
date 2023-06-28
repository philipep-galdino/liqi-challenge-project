package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: I don't want %v, I want %v", status, http.StatusOK)
	}

	expected := "Request sent to: /\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGenerateKeysHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/keys", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(generateKeys)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: I don't want %v, I want %v",
			status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	if !strings.Contains(responseBody, "Private key:") || !strings.Contains(responseBody, "Public key:") {
		t.Errorf("handler returned unexpected body: %v", responseBody)
	}
}
