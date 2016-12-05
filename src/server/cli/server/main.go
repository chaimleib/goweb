package main

import (
  "server"
  "os"
  "fmt"
  "log"
)

var Log *log.Logger

func setup() {
  Log = log.New(os.Stderr, "", log.Ltime | log.Lshortfile)
  Log.Print("Log initialized")
}

func main() {
  port := "8080"
  setup()
  Log.Print("Running goweb at http://localhost:" + port)
  fmt.Println("Use ctrl-c to exit")
  server.Start(port)
}
