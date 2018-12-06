package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := int64(30)
	x := Sum(1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4)
	if x != result {
		t.Error("Expected", result, "Got", x)
	}
}

func TestAvg(t *testing.T) {
	result := 4.5
	x := Avg(8, 1)
	if x != result {
		t.Error("Expected", result, "Got", x)
	}
}
