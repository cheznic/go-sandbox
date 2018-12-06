package main

func main() {

}

// Sum accepts numerous interger arguments and return the sum of those integers.
func Sum(xi ...int64) int64 {
	var sum, v int64
	for _, v = range xi {
		sum += v
	}
	return sum
}

// Avg accepts multiple integer arguments and returns the average of those integers.
func Avg(xi ...int64) float64 {
	var i int
	var sum, v int64
	for i, v = range xi {
		sum += v
	}
	return float64(sum) / float64(i+1)
}
