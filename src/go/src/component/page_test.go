package component

import "testing"

func check(e error, t *testing.T) {
  if e != nil {
    t.Error(e)
  }
}

func TestPage(t *testing.T) {
  body := "This is a test.\n\nHello world!"
  p1 := &Page{Slug: "testdata/hello", Body: body}
  p1.Save()
  p2, err := LoadPage("testdata/hello")
  check(err, t)
  if body != p2.Body {
    t.Error("Body as written != Body read")
  }
}
