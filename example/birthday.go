package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	// type Config struct {
	// 	Name     string
	// 	Birthday string
	// 	Twitter  string
	// 	Location string

	// 	type Config struct {
	// 	Name string
	// 	Birthday string
	// 	Twitter string
	// 	Location string
	// 	Pets []struct {
	// 		Species string
	// 		PetName string `json:"pet-name"`
	// 	}
	// }

	type LittleMonster struct {
		Species string
		PetName string `json:"pet-name"`
	}

	type Config struct {
		Name     string
		Birthday string
		Twitter  string
		Location string
		Pets     []LittleMonster
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
	fmt.Println(config.Name, config.Birthday, config.Twitter, config.Location) // output
}
