package main

import (
  "fmt"
  "net/http"
)

func HelloWorld(response http.ResponseWriter, request *http.Request) {
  fmt.Fprint(response, "Hello World!")
}

func main() {
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(":3000", nil)
}
