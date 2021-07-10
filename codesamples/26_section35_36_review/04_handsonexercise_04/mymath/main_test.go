package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}
	te := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}
	for _, value := range te {
		if CenteredAvg(value.data) != value.result {
			t.Error("Expected :", value.result, "Got :", value.result)

		}
	}

}
func ExampleCenteredAvg() {
	type test struct {
		data   []int
		result float64
	}
	te := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}
	for _, value := range te {
		fmt.Println(CenteredAvg(value.data))
		//Output:
		//6
		//9
		//9
		//170

	}

}
func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}

}
