package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/test", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "test")
}

