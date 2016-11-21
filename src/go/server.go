package main

import (
  "os"
  "net/http"
  "html/template"
  "components"
)

type Root struct {
  Title string
  Header string
  Body string
  Footer string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  p := r.URL.Path
  tmplData, err := loadPage("../go-templates/root.html")
  if (err != nil) { panic(err) }
  tmpl, err := template.New("root").Parse(tmplData.Body)
  if (err != nil) { panic(err) }
  pageData := Root{tmplData.Slug, "Header<br/>", "Hi there!", "Footer"}
  err = tmpl.Execute(os.Stdout, pageData)
  if (err != nil) { panic(err) }
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}

