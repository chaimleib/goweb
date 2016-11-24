package server

import (
  "html/template"
  "sync"
  "server/component"
)

type templateSet struct {
  mu sync.Mutex
  set map[string]*template.Template
}

// Where all the templates get loaded first time they are requested
var ts = &templateSet{
  set: make(map[string]*template.Template),
}

func getTemplate(name string) (*template.Template, error) {
  ts.mu.Lock()
  defer ts.mu.Unlock()
  if tmpl, ok := ts.set[name]; ok {
    return tmpl, nil
  }
  tmpl, err := _loadTemplate(name)
  return tmpl, err
}

func setTemplate(name string) error {
  tmpl, err := _loadTemplate(name)
  if err != nil {
    return err
  }
  ts.mu.Lock()
  defer ts.mu.Unlock()
  ts.set[name] = tmpl
  return nil
}

func _loadTemplate(name string) (*template.Template, error) {
  tmplData, err := component.LoadPage("go-templates", name)
  if err != nil {
    return nil, err
  }
  tmpl, err := template.New(name).Parse(tmplData.Body)
  return tmpl, err
}
