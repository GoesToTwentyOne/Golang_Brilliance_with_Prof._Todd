package main

import (
	"fmt"
)

func foo() {
	fmt.Println("I'm from foo")
}
func bar(s string) {
	fmt.Println("Hi! I'm ", s)
}
func woor(s string) string {
	return fmt.Sprintln("I'm from from wooor,", s)
}
func harto(x int, s string) (string, bool) {
	r1 := fmt.Sprintln("age is ", x, `"I'm Md. Nihad Hossain"`)
	r2 := true
	return r1, r2
}
func main() {
	foo()
	bar("Md. Nihad Hossain")
	s := woor("Alex Goot")
	fmt.Println(s)
	b, s2 := harto(21, "Hena")
	fmt.Println(b)
	fmt.Println(s2)

}
