package main

import (
    "fmt"
    "net/http"
)

// handleRequest processes an HTTP request and writes a response.
func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Request handled")
}

// serve is the HTTP handler registered with the server. It forwards the request
// to handleRequest. The original code called a miss‑spelled identifier
// `handleReqest`, which caused a compilation error. The call is now corrected
// to use the proper function name and matches the function signature.
func serve(w http.ResponseWriter, r *http.Request) {
    handleRequest(w, r)
}

func main() {
    http.HandleFunc("/", serve)
    http.ListenAndServe(":8080", nil)
}
