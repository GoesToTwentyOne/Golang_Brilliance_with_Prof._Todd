package mymath

import (
	"fmt"
	"testing"
)

type tabtest struct {
	data   []int
	result int
}

var ta = []tabtest{
	{data: []int{1, 4, 6, 8, 100}, result: 6},
	{data: []int{0, 8, 10, 1000}, result: 9},
	{data: []int{9000, 4, 10, 8, 6, 12}, result: 9},
	{data: []int{123, 744, 140, 200}, result: 170},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range ta {
		if CenteredAvg(v.data) != float64(v.result) {
			t.Error("Expected :", v.result, "Got : ", CenteredAvg(v.data))
		}
	}

}

func ExampleCenteredAvg() {
	for _, v := range ta {
		fmt.Println(CenteredAvg(v.data))
	}
	//Output:
	//6
	//9
	//9
	//170
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range ta {
			CenteredAvg(v.data)

		}
	}
}
