package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("Server starting on port 32777...")
    http.ListenAndServe(":32777", nil)
}
