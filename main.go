package main

import (
  "net/http"
  "github.com/russross/blackfriday"
)

/*
  Generate the markdown that has been entered into the form. This markdown
  should be generated into HTML, which is viewable once the user hits the
  'Submit' button on the form
*/
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
  markdown := blackfriday.MarkdownCommon([]byte (r.FormValue("body")))
  rw.Write(markdown)
}

func main() {
  // http.ListenAndServe(":8080", http.FileServer(http.Dir("./public")))

  // Serve the markdown route
  http.HandleFunc("/markdown", GenerateMarkdown)
  http.Handle("/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":8080", nil)
}
