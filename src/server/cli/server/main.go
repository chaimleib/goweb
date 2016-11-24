package main

import (
  "server"
  "fmt"
)

func main() {
  port := "8080"
  fmt.Println("Running goweb at http://localhost:" + port)
  fmt.Println("Use ctrl-c to exit")
  server.Start(port)
}
