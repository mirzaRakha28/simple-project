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
    fmt.Println("Server started at :9000")
    if err := http.ListenAndServe(":9000", nil); err != nil {
        fmt.Println(err)
    }
}

