package server

import (
  "html/template"
  "sync"
  "server/component"
  "errors"
)

type templateCache struct {
  mu sync.Mutex
  data map[string]*template.Template
}

// Templates get loaded from the filesystem into `templates` first time they are
// requested. This cache minimizes the number of filesystem hits, as well as the
// number of times we have to parse the templates.
var templates = &templateCache{
  data: make(map[string]*template.Template),
}

// Location of the HTML template files
const TemplatesDir = "go-templates"

// Given the name of a template file in TemplatesDir, returns its compiled
// template. If in the cache, returns the cached template. If not in the cache
// yet, loads it from the file system into the cache before returning it. If the
// template file cannot be found, returns nil and an error.
// The `name` argument should specify the basename of the file, minus the
// ".html" extension.
func (self *templateCache) Get(name string) (*template.Template, error) {
  if tmpl := self.GetCache(name); tmpl != nil {
    return tmpl, nil
  }
  if err := self.Load(name); err != nil {
    return nil, err
  }
  tmpl := self.GetCache(name)
  if (tmpl != nil) {
    return tmpl, nil
  }
  return nil, errors.New("failed to load file into templateCache")
}

// (Re)loads the named template file from TemplatesDir into the cache. The `name`
// argument should specify the basename of the file, minus the ".html"
// extension.
func (self *templateCache) Load(name string) error {
  tmplData, err := component.LoadPage(TemplatesDir, name)
  if err != nil {
    return err
  }
  tmpl, err := template.New(name).Parse(tmplData.Body)
  if err != nil {
    return err
  }
  self.SetCache(name, tmpl)
  return nil
}

// Checks our cache for `name` returns it, if present. If not in the cache yet,
// returns nil.
// The `name` argument should specify the basename of the file, minus the
// ".html" extension.
func (self *templateCache) GetCache(name string) *template.Template {
  self.mu.Lock()
  defer self.mu.Unlock()
  if tmpl, ok := self.data[name]; ok {
    return tmpl
  }
  return nil
}

// Stores `tmpl` as `name`. If an entry already exists for `name`, overwrite it.
func (self *templateCache) SetCache(name string, tmpl *template.Template) {
  self.mu.Lock()
  defer self.mu.Unlock()
  self.data[name] = tmpl
}
