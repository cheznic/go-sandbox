package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	type input struct {
		data   []int64
		result int64
	}

	inputs := []input{
		input{
			data:   []int64{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
			result: 30,
		},
		input{
			data:   []int64{-1, -2, -3, -4, -1, -2, -3, -4, -1, -2, -3, -4},
			result: -30,
		},
		input{
			data:   []int64{8},
			result: 8,
		},
		input{
			data:   []int64{555555555, 111111111, 111111111, 111111111, 111111111, 1},
			result: 1000000000,
		},
	}

	for _, v := range inputs {
		result := Sum(v.data...)
		if result != v.result {
			t.Error("Expected", v.result, "Got", result)
		}
	}
}

func TestAvg(t *testing.T) {

	type input struct {
		data   []int64
		result float64
	}

	inputs := []input{
		input{
			data:   []int64{8, 1},
			result: 4.5,
		},
		input{
			data:   []int64{100000000, 0},
			result: 50000000,
		},
		input{
			data:   []int64{-3, -3, -3, -3, -3, -3, 3, 3, 3, 3, 3, 3},
			result: 0,
		},
	}

	for _, v := range inputs {
		result := Avg(v.data...)
		if result != v.result {
			t.Error("Expected", result, "Got", result)
		}
	}
}
