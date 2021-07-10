package main

import (
	"fmt"
	"testing"
)

func TestFindOutlier(t *testing.T) {
	if x := FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}); x != 160 {
		t.Error("I expected : 160 Got: ", x)
	}

}
func ExampleFindOutlier() {
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	//Output:
	//160

}
func BenchmarkFindOutlier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21})

	}
}
