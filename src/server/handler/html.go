package handler

import (
  "net/http"
  "html/template"
)

type PageContext interface{}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func HTML(w http.ResponseWriter, t *template.Template, c PageContext) {
  err := t.Execute(w, c)
  check(err)
}
