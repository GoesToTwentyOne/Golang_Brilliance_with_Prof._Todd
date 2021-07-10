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
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}
	for _, value := range te {
		f := CenteredAvg(value.data)
		if f != value.result {
			t.Errorf("expected : %v got : %v", value.result, value.result)

		}

	}

}
func ExampleCenteredAvg() {
	type test struct {
		data   []int
		result float64
	}

	te := []test{
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}
	for _, value := range te {

		fmt.Println(CenteredAvg(value.data))
		//Output:
		//40
		//7
		//5

	}

}
func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}
