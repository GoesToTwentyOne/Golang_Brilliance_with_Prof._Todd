package bencho

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	x := Greeting("Nihad Hossain")
	if x != "MY name is Nihad Hossain " {
		t.Error("I expected: My name is Nihad Hossain Got:", x)

	}

}
func ExampleTestGreeting() {
	fmt.Println(Greeting("Nihad Hossain"))
	//Output:
	//MY name is Nihad Hossain

}
func BenchmarkGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("James")
	}
}
