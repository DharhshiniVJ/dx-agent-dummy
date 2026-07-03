package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Fixed typo: call the correctly named function
        handleRequest(w, r)
    })
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}

// handleRequest processes an incoming HTTP request.
func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
