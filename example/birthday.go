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

const ctLayout = "January 2, 2006"

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
		Birthday CustomTime
		Twitter  string
		Location string
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
	fmt.Println(config.Name, config.Birthday.Format(ctLayout), config.Twitter, config.Location) // output
}
