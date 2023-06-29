package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/sha3"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request sent to: %s\n", r.URL.Path)
}

func generateKeys(w http.ResponseWriter, _ *http.Request) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		http.Error(w, "Error while generating keys", http.StatusInternalServerError)
		return
	}

	publicKey := privateKey.PublicKey

	privateKeyBytes := privateKey.D.Bytes()
	publicKeyBytes := elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y)

	fmt.Fprintf(w, "Private key: %s\n", hex.EncodeToString(privateKeyBytes))
	fmt.Fprintf(w, "Public key: %s\n", hex.EncodeToString(publicKeyBytes))

}

func getAddress(w http.ResponseWriter, r *http.Request) {
	publicKeyHex := r.URL.Query().Get("publicKey")
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)

	if err != nil {
		http.Error(w, "Invalid public key", http.StatusBadRequest)
		return
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	address := hash.Sum(nil)[12:]

	fmt.Fprintf(w, "Ethereum Address: %s\n", hex.EncodeToString(address))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/keys", generateKeys)
	http.HandleFunc("/address", getAddress)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
