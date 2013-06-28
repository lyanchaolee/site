package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Coming soon......")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

