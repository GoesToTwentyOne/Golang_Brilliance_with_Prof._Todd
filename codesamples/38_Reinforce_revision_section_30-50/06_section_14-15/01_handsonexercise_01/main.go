package main

import "fmt"

func main() {
	f := foo()
	fmt.Println("my age is ", f)
	b, s := bar()
	fmt.Println(b, s)

}
func foo() int {
	return 21
}
func bar() (int, string) {
	return 21, "Hello from bar "

}
