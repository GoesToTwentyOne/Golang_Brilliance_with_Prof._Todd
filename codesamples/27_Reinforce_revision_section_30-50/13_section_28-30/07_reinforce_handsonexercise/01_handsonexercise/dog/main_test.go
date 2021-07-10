package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	if Years(10) != 70 {
		t.Error("Expected : 70 Got :", Years(10))
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	// Output :
	// 70

}
func TestYearsTwo(t *testing.T) {
	if YearsTwo(20) != 140 {
		t.Error("Expected : 140 Got :", YearsTwo(20))
	}
}
func ExampleYearsTwo() {
	fmt.Println(Years(20))
	// Output :
	// 140

}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}

}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}

}
