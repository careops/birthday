package main

import (
	"bufio"
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

// TODO: support short date format in addition to long, e.g. "Jan 2"
const readableBirthdayLayout = "January 2"

func (ct *ReadableBirthday) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(readableBirthdayLayout, s)
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

	// TODO: display birthdays this week

	// the 4 types of for loops in go:
	//
	// infinite
	// for {
	//   ...
	//}
	//
	// range
	// for i, v := range configs {
	//   ...
	//}
	//
	// c-like
	// for i := 0; i < n; i++ {
	//   ...
	//}
	//
	// while
	// for i < n {
	//   ...
	//}

	// func (*Reader) ReadLine
	// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

	// line = "yes"
	// [121 101 115]

	acceptedTermsAndConditions := false
	for acceptedTermsAndConditions == false {
		fmt.Print("are you willing to provide your soul to the devil? ")
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		//if len(line) == 3 && line[0] == 121 && line[1] == 101 && line[2] == 115 {
		if string(line) == "yes" {
			acceptedTermsAndConditions = true
			fmt.Println("i have taken your soul, mortal!")
		}
		// fmt.Println(line)
		// fmt.Println(string(line))
	}

	//for {
	for _, v := range configs {
		//if v.Birthday.Format(readableBirthdayLayout) == #ISOweek
		//{
		fmt.Println(v.Name, v.Birthday.Format(readableBirthdayLayout), v.Twitter, v.Location) // output
	}
	//}

}
