package main

import "net/http"

func serve() {
	handleRequest(nil, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Handle it
}
