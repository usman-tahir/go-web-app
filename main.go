package main

import (
  "log"
  "net/http"
  "github.com/codegangsta/negroni"
)

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  log.Println("[Begin] Logging")
  if r.URL.Query().Get("password") == "password" {
    next(rw, r)
    log.Println("Access granted.")
  } else {
    http.Error(rw, "Error: Access denied.", 401)
    log.Println("Access denied.")
  }
  log.Println("[End] Logging")
}

func main() {
  // Middleware stack
  n := negroni.New(
    negroni.NewRecovery(),
    negroni.HandlerFunc(MyMiddleware),
    negroni.NewLogger(),
    negroni.NewStatic(http.Dir("public")),
  )

  n.Run(":3000")
}
