package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	if Years(10) != 70 {
		t.Error("expected : 70 got: ", Years(10))
	}

}
func TestYearsTwo(t *testing.T) {
	if YearsTwo(10) != 70 {
		t.Error("expected : 70 got: ", YearsTwo(10))

	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70

}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
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
