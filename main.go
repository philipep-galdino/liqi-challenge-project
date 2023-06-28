package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request sent to: %s\n", r.URL.Path)
}

func generateKeys(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/keys", generateKeys)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
