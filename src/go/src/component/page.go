package component

import (
  "io/ioutil"
)

type Page struct {
  Dir string
  Slug string
  Body string
}

const dtype = "html"

func (p *Page) Save() error {
  filename := p.Dir + "/" + p.Slug + "." + dtype
  return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func LoadPage(dir string, slug string) (*Page, error) {
  filename := dir + "/" + slug + "." + dtype
  bodyBytes, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  body := string(bodyBytes)
  return &Page{Dir: dir, Slug: slug, Body: body}, nil
}
