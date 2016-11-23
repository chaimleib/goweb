package main

import (
  "net/http"
  "html/template"
  "component"
)

type Root struct {
  Title string
  Header string
  Body string
  Footer string
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  // p := r.URL.Path
  tmplData, err := component.LoadPage("../go-templates/root")
  check(err)
  tmpl, err := template.New("root").Parse(tmplData.Body)
  check(err)
  pageData := Root{tmplData.Slug, "Header", "Hi there!", "Footer"}
  err = tmpl.Execute(w, pageData)
  check(err)
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
