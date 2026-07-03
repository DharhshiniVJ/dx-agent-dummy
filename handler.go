package main

import "net/http"

func serve() {
	handleReqest(nil, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Handle it
}
