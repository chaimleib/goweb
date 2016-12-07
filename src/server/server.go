package server

import (
  "net/http"
  "server/handler"
  "fmt"
  "server/config"
  "log"
)

var Log *log.Logger

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

func Start(cfg *config.ConfigObject, log *log.Logger) {
  tcpAddr := fmt.Sprintf(":%d", cfg.Port)
  log.Printf("Running server at %s%s", cfg.Host, tcpAddr)
  http.HandleFunc("/", testHandler)
  log.Fatal(http.ListenAndServe(tcpAddr, nil))
}
