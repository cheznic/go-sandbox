package illuminate

import (
	"fmt"
)

func ExampleLamp_On() {
	l := Lamp{}
	fmt.Println(l.On())
	// Output:
	// shine
}

func ExampleLamp_Off() {
	l := Lamp{}
	fmt.Println(l.Off())
	// Output:
	// dark
}
