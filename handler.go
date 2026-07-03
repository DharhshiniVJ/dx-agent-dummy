package main

import "fmt"

func handleRequest(id int) {
    fmt.Println("Request", id)
}

func main() {
    fmt.Println("Hello World")
    // Fixed typo: corrected function name from handleReqest to handleRequest
    handleRequest(42)
}
