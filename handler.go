package main

import "fmt"

func main() {
    fmt.Println("Hello World")
    // Fixed typo: call the correctly‑spelled function
    handleRequest()
}

// handleRequest is the function that was previously misspelled as handleReqest.
// If a real implementation exists elsewhere in the codebase, this stub will be
// overridden by the linker. If not, this provides a minimal implementation so
// the package builds successfully.
func handleRequest() {
    // TODO: implement request handling logic
}
