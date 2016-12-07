package main

import (
  "os"
  "log"
  "fmt"
  "server"
  "server/config"
)

var Log *log.Logger

func setup() *config.ConfigObject {
  Log = log.New(os.Stderr, "", log.Ltime | log.Lshortfile)
  Log.Print("Log initialized")
  cfg := &config.ConfigObject{
    Host: "http://localhost",
    Port: 8080,
    Environment: config.EnvDev,
  }
  Log.Printf("Setting options: %+v\n", cfg)
  return cfg
}

func main() {
  cfg := setup()
  fmt.Printf("\nStarting server, use ctrl-c to exit\n\n")
  server.Start(cfg, Log)
}
