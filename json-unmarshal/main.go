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
	js := `
	[
		{
			"First":"James",
			"Last":"Bond",
			"Age":32,
			"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]
		},
		{
			"First":"Miss",
			"Last":"Moneypenny",
			"Age":27,
			"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]
		},
		{
			"First":"M",
			"Last":"Hmm",
			"Age":54,
			"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]
		}
	]`

	var people []Person

	if err := json.Unmarshal([]byte(js), &people); err != nil {
		fmt.Println("error", err)
	}

	for i, person := range people {
		fmt.Println("Person "+fmt.Sprint(i)+":", person.First, person.Last+",", "Age:", person.Age)
		for ii, saying := range person.Sayings {
			fmt.Println("\tSaying "+fmt.Sprint(ii)+":", saying)
		}
	}
}
