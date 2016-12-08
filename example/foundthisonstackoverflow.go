package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	type Config struct {
		name     []string
		birthday []string
		twitter  []string
		location []string
	}

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(config.name) // output
}
