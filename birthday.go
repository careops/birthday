package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	filepath := ""
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
	//TODO birthday-test.go to check for edge cases around ISOWeek, end of year

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The following people have birthdays this week:*")
		birthdayCount := 0

		for _, v := range configs {
			_, week := v.Birthday.ISOWeek()
			_, currentWeek := time.Now().ISOWeek()
			if week == currentWeek {

				fmt.Fprintf(w, "\n %v, on %v 🎂 🎈 🎉\n", v.Name, v.Birthday.Format(readableBirthdayLayout))
				birthdayCount++
			}

		}
		if birthdayCount == 0 {
			fmt.Fprintf(w, "Sorry, no birthdays this week, check back next week")
		}
	})
	fmt.Println("Listening and Serving on localhost:8080/  Go there!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
