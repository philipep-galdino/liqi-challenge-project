package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/philipep-galdino/liqi-challenge-project/pkg/ethereum"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, wish to get %v", status, http.StatusOK)
	}

	expected := "Request sent to: /\n"

	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

func TestGenerateKeys(t *testing.T) {
	req, err := http.NewRequest("GET", "/keys", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateKeys)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	if !strings.Contains(responseBody, "Private key:") || !strings.Contains(responseBody, "Public key:") {
		t.Errorf("unexpected body: %v", responseBody)
	}
}

func TestGetAddress(t *testing.T) {
	_, publicKey, _ := ethereum.GenerateKeys()

	req, err := http.NewRequest("GET", "/address?publicKey="+publicKey, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAddress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code returned: got %v, wanted %v", status, http.StatusOK)
	}

	responseBody := rr.Body.String()

	if !strings.Contains(responseBody, "Ethereum Address:") {
		t.Errorf("got unexpected body: %v", responseBody)
	}
}
