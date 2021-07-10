package verocell

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2, 4, 3) != 10 {
		t.Error("expected : 10 Got :", Add(1, 2, 4, 3))
	}

}
func ExampleAdd() {
	fmt.Println(Add(1, 2, 4, 3))
	//Output:
	//10
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2, 4, 3)

	}

}
