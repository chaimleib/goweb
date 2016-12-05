package server

import (
  "net/http"
  "server/handler"
)

type Context struct {
  Title string
  Header string
  Body string
  Footer string
  CSS []string
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  // p := r.URL.Path
  css := []string{"test"}
  pageData := Context{
    Title: "Test Root",
    Header: "Header",
    Body: "Hi there!",
    Footer: "Footer",
    CSS: css,
  }
  t, err := templates.Get("root")
  check(err)
  handler.HTML(w, t, pageData)
}

func Start(port string) {
  http.HandleFunc("/", testHandler)
  http.ListenAndServe(":" + port, nil)
}
