package main

import "fmt"

func main() {
	defer foo()
	bar()

}
func foo() {
	fmt.Println("Foo from foo")
}
func bar() {
	fmt.Println("Foo from bar")
}
