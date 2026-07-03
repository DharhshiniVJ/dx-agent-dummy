package main

import (
    "fmt"
    "net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Handled")
}

// Handler is the HTTP handler used in main.
func Handler(w http.ResponseWriter, r *http.Request) {
    // Some processing...
    handleRequest(w, r) // corrected typo
}

func main() {
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":8080", nil)
}
