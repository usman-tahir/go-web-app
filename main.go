package main

import (
  "net/http"
  "./src/gopkg.in/unrolled/render.v1"
)

func main() {
  r := render.New(render.Options{})
  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Welcome! Feel free to browse the other pages."))
  })

  mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
    r.Data(w, http.StatusOK, []byte("Binary data"))
  })

  mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
    r.JSON(w, http.StatusOK, map[string]string{"Hello": "JSON"})
  })

  mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
    r.HTML(w, http.StatusOK, "Exxamp", nil)
  })

  http.ListenAndServe(":3000", mux)
}
