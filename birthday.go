package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type ReadableBirthday struct {
	time.Time
}

const ctLayout = "January 2"

func (ct *ReadableBirthday) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func main() {

	type Config struct {
		Name     string
		Birthday ReadableBirthday
		Twitter  string
		Location string
	}

	if len(os.Args) == 1 {
		fmt.Println("I need a .json file to read from, please add this to your go run birthday.go command.")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("You gave me too many .json files to read from, please only use one .json in your go run birthday.go command.")
		os.Exit(1)
	}

	var filepath string
	filepath = os.Args[1]

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	configs := []Config{}
	err = decoder.Decode(&configs)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	//loops around the code address each element of list using a variable

	fmt.Println(configs[0].Name, configs[0].Birthday.Format(ctLayout), configs[0].Twitter, configs[0].Location) // output
}