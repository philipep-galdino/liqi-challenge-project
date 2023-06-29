package http

import (
	"fmt"
	"net/http"

	"github.com/philipep-galdino/liqi-challenge-project/pkg/ethereum"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request sent to: %s\n", r.URL.Path)
}

func GenerateKeys(w http.ResponseWriter, _ *http.Request) {
	privateKey, publicKey, err := ethereum.GenerateKeys()

	if err != nil {
		http.Error(w, "Error while generating keys", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Private key: %s\n", privateKey)
	fmt.Fprintf(w, "Public key: %s\n", publicKey)
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	publicKey := r.URL.Query().Get("publicKey")
	address, err := ethereum.GetAddress(publicKey)

	if err != nil {
		http.Error(w, "Invalid public key", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Ethereum Address: %s\n", address)
}
