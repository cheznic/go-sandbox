package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge .
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast .
type ByLast []user

func (user ByLast) Len() int           { return len(user) }
func (user ByLast) Swap(i, j int)      { user[i], user[j] = user[j], user[i] }
func (user ByLast) Less(i, j int) bool { return user[i].Last < user[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	sort.Sort(ByAge(users))

	fmt.Println("Sorted by Age:")

	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println()

	sort.Sort(ByLast(users))

	fmt.Println("Sorted by Last:")

	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println()

	var sayings []string

	for _, user := range users {

		sayings = user.Sayings
		sort.Strings(sayings)
		fmt.Println(user.First, user.Last+"'s", "Sayings")
		for ii, saying := range user.Sayings {
			fmt.Println("Saying "+fmt.Sprint(ii+1)+":", saying)
		}
		fmt.Println()

	}

}
