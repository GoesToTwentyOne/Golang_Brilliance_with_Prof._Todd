package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		result int
	}
	te := []test{
		//readability practice
		//test{[]int{10, 3, -1, 2}, 14},
		{[]int{1, 3, -1, 2}, 5},
		{[]int{1, 10, -1, 2}, 12},
		{[]int{1, 13, -1, 2}, 15},
		{[]int{1, 23, -4, 2}, 22},
		{[]int{4, 3, -1, 2}, 8},
		{[]int{6, 3, -1, 2}, 10},
		{[]int{10, 3, -1, 2}, 14},
	}
	for _, value := range te {
		x := sum(value.data...)
		if x != value.result {
			t.Error("I expected", value.result, " got ", x)

		}

	}

}
