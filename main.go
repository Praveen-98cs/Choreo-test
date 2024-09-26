package main

import (
    "fmt"
    "net/http"
    "time"
)

const timeout = 5 // Hardcoded timeout in seconds

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Wait for the hardcoded timeout
    time.Sleep(time.Duration(timeout) * time.Second)

    // Send the response
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Printf("Server is running on http://localhost:8080 with a %d second timeout\n", timeout)
    http.ListenAndServe(":8080", nil)
}
