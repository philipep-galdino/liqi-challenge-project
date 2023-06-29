package main

import (
	"fmt"
	"log"
	"net/http"

	customHTTP "github.com/philipep-galdino/liqi-challenge-project/pkg/http"
)

func main() {
	http.HandleFunc("/", customHTTP.Handler)
	http.HandleFunc("/keys", customHTTP.GenerateKeys)
	http.HandleFunc("/address", customHTTP.GetAddress)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
