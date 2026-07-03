package main

import (
    "fmt"
    "net/http"
)

// handleRequest processes an incoming HTTP request and writes a simple response.
func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    // Fixed typo: was `handleReqest`, now correctly references `handleRequest`.
    http.HandleFunc("/", handleRequest)
    // Start the HTTP server; any error is logged and the program exits.
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("server error: %v\n", err)
    }
}
