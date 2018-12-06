package main

import (
	"encoding/json"
	"fmt"
)

// Person .
type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	p1 := Person{
		First:   "James",
		Last:    "Bond",
		Age:     32,
		Sayings: []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In his majesty's royal service"},
	}

	p2 := Person{
		First:   "Miss",
		Last:    "Moneypenny",
		Age:     27,
		Sayings: []string{"James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret agent myself."},
	}

	p3 := Person{
		First:   "M",
		Last:    "Hmm",
		Age:     54,
		Sayings: []string{"Oh, James. You didn't.", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"},
	}

	people := []Person{p1, p2, p3}

	var js []byte
	var err error

	if js, err = json.Marshal(&people); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(string(js))
}
