package main

import (
	"fmt"
	"net/http"
)

func main() {
    // Define a handler function for your HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Set the response header
        w.Header().Set("Content-Type", "text/plain")

        // Write the response content
        fmt.Fprintln(w, "Hello, World!")
    })

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error:", err)
    }
}

