package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Println(rw, "Home\n")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Println(rw, "Posts Index\n")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Println("Posts Create\n")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := p.ByName("id")
  fmt.Fprintln(rw, "Showing post", id, "\n")
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "Post Update\n")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "Post Delete\n")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "Post Edit\n")
}

func main() {
  r := httprouter.New()
  r.GET("/", HomeHandler)

  // Posts collection
  r.GET("/posts", PostsIndexHandler)
  r.POST("/posts", PostsCreateHandler)

  // Single post handling
  r.GET("/posts/:id", PostShowHandler)
  r.PUT("/posts/:id", PostUpdateHandler)
  r.GET("/posts/:id/edit", PostEditHandler)

  fmt.Println("Starting server on localhost:8080")
  http.ListenAndServe(":8080", r)
}
