package main

import (
    "fmt"
    "net/http"
)

// handleRequest processes an HTTP request and writes a response.
func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "request handled")
}

func main() {
    http.HandleFunc("/", handleRequest) // corrected typo: handleReqest -> handleRequest
    fmt.Println("Server listening on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("server error: %v\n", err)
    }
}
