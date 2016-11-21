package components

import (
  "fmt"
  "io/ioutil"
)

type Page struct {
  Slug string
  Body []byte
}

const dtype = "html"

func (p *Page) save() error {
  filename := p.Slug + "." + "html"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(slug string) (*Page, error) {
  filename := slug + "." + "html"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Slug: slug, Body: body}, nil
}

func main() {
  p1 := &Page{Slug: "hello", Body: []byte("This is a test.\n\nHello world!")}
  p1.save()
  p2, _ := loadPage("hello")
  fmt.Println(string(p2.Body))
}

