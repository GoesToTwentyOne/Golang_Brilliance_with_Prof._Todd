package main

import (
	"fmt"
)

func foo() {
	fmt.Println("I'm from foo")
}
func bar() {
	fmt.Println("I'm from bar")
}
func main() {
	defer foo()
	bar()

}
