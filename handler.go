package main

import "fmt"

func main() {
    fmt.Println("Hello World")
    // Fixed typo: call the correctly‑spelled function
    handleRequest()
}

// Added the missing function definition. The original code attempted to call
// `handleReqest`, which does not exist. Providing the correctly‑named function
// resolves the undefined identifier error.
func handleRequest() {
    fmt.Println("Handling request")
}
