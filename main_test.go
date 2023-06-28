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

func TestGetAddressHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/address?publicKey=04b86a5bd672fdc23458b4e6fbaa2638f250bd7fb784629563c713bb9b8277cbcfb8f1186160c6a10c39ab6262a199e1082c6f307398c5e132fb596e2a635e4a27", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAddress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status: don't want %v, I want %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	if !strings.Contains(responseBody, "Ethereum Address:") {
		t.Errorf("handler returned unexpected body: %v", responseBody)
	}
}
