package cat

import (
	"fmt"
	"strings"
	"testing"
)

func TestMystring(t *testing.T) {
	xs := strings.Split("My name is Nihad", " ")

	if s := Mystring(xs); s != Mystring(xs) {
		t.Errorf("expected My name is Nihad \n got %s ", Mystring(xs))

	}

}
func TestJoin(t *testing.T) {
	xs := strings.Split("My name is Nihad ", " ")
	if Join(xs) != "My name is Nihad " {
		t.Errorf("expected My name is Nihad \n got %s ", Join(xs))

	}
}
func ExampleMystring() {
	xs := strings.Split("My name is Nihad ", " ")
	fmt.Println(Mystring(xs))
	//Output:
	//My name is Nihad

}
func ExampleJoin() {
	xs := strings.Split("My name is Nihad ", " ")
	fmt.Println(Join(xs))
	//Output:
	//My name is Nihad

}
func BenchmarkMystring(b *testing.B) {
	xs := strings.Split("My name is Nihad ", " ")
	for i := 0; i < b.N; i++ {
		Mystring(xs)

	}

}
func BenchmarkJoin(b *testing.B) {
	xs := strings.Split("My name is Nihad ", " ")
	for i := 0; i < b.N; i++ {
		Join(xs)

	}

}
