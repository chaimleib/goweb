package component

import (
  "io/ioutil"
)

type Page struct {
  Slug string
  Body string
}

const dtype = "html"

func (p *Page) Save() error {
  filename := p.Slug + "." + "html"
  return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func LoadPage(slug string) (*Page, error) {
  filename := slug + "." + "html"
  bodyBytes, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  body := string(bodyBytes)
  return &Page{Slug: slug, Body: body}, nil
}
