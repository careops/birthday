package main

import (
    "encoding/json"
    "os"
    "fmt"
)

type Config struct {
    name    []string
    birthday []string
    twitter []string
    location []string
}

file, _ := os.Open("conf.json")
decoder := json.NewDecoder(file)
config := Config{}
err := decoder.Decode(&configuration)
if err != nil {
  fmt.Println("error:", err)
}
fmt.Println(config.name) // output
