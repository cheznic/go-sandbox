package main

import (
	"fmt"
)

func main() {

	//
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])

	if v, ok := m["Chas"]; ok {
		fmt.Println("First try to print Chas:", v)
	}

	m["Chas"] = 52

	if v, ok := m["Chas"]; ok {
		fmt.Println("Second try to print Chas:", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Chas")

	for k, v := range m {
		fmt.Println(k, v)
	}

}
