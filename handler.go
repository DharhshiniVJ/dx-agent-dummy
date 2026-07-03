package main

import (
    "fmt"
    "net/http"
    // other imports
)

// ... other code ...

func someHandler(w http.ResponseWriter, r *http.Request) {
    // ... some logic ...
    // line 31 (originally): handleReqest(w, r)
    // corrected spelling:
    handleRequest(w, r)
    // ... rest of function ...
}

// ... rest of file unchanged ...
