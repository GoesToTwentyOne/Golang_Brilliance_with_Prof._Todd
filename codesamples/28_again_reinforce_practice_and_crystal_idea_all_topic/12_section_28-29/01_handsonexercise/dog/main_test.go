package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	if x := Years(10); x != 70 {
		t.Error("expected 70 got ", x)
	}

}
func TestYearsTwo(t *testing.T) {
	if x := YearsTwo(20); x != 140 {
		t.Error("expected 140 got ", x)
	}

}
func ExampleYears() {
	fmt.Println(Years(10))
	//output:
	//70

}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	//output:
	//140

}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
