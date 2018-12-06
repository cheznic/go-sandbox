// Package numath contains functions that do math things
package numath

// Sum returns the sum of all integers passed in.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
