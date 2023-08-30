package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
  fmt.Printf("HTTPサーバ立ち上げ")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}