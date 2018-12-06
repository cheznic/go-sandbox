// Package saying provides some convient sayings
package saying

import "fmt"

// Greet accepts a string and returns a greeting with that string value in it.
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}
