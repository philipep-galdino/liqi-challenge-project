package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request sent to: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
