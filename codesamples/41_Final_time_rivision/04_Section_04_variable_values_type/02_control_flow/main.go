package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi I'm from First line")
	foo()
	bar()
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
	fmt.Println("I'm from last line")

}
func foo() {
	fmt.Println("I'm from FOO")
}
func bar() {
	fmt.Println("I'm form BAR")
}
