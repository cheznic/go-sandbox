package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

// Address todo
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Person todo
type Person struct {
	ID             string
	First          string
	Last           string
	MailingAddress Address
}

func main() {

	address1 := Address{
		Street: "123 Lincoln Ln",
		City:   "Walla Walla",
		State:  "WA",
		Zip:    "98656",
	}

	person1 := Person{
		ID:             uuid.New().String(),
		First:          "Jon",
		Last:           "Stuart",
		MailingAddress: address1,
	}

	address2 := Address{
		Street: "1600 Pennsylvania Ave NW",
		City:   "Washington",
		State:  "DC",
		Zip:    "20006",
	}

	person2 := Person{
		ID:             uuid.New().String(),
		First:          "Bill",
		Last:           "Maher",
		MailingAddress: address2,
	}

	people := map[string]Person{
		person1.ID: person1,
		person2.ID: person2,
	}

	bytes, err := json.MarshalIndent(people, "", "   ")
	//bytes, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	n, err := os.Stdout.Write(bytes)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Bytes written: ", n)
	fmt.Println()
	fmt.Println()

	unmarshal(bytes)
}

func unmarshal(jsonBlob []byte) {

	var people map[string]Person

	json.Unmarshal(bytes, &people)

	fmt.Println(people)
}

var bytes = []byte(`{
	"73b0d6dd-1603-4918-9555-1b34ca16d302": {
		"ID": "73b0d6dd-1603-4918-9555-1b34ca16d302",
		"First": "Bill",
		"Last": "Maher",
		"MailingAddress": {
			"Street": "1600 Pennsylvania Ave NW",
			"City": "Washington",
			"State": "DC",
			"Zip": "20006"
		}
	},
	"7cbc5734-ebe1-437d-b018-e9c11af75396": {
		"ID": "7cbc5734-ebe1-437d-b018-e9c11af75396",
		"First": "Jon",
		"Last": "Stuart",
		"MailingAddress": {
			"Street": "123 Lincoln Ln",
			"City": "Walla Walla",
			"State": "WA",
			"Zip": "98656"
		}
	}
}`)