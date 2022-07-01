package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

}
func foo() int {
	x := 43
	return x

}
func bar() (int, string) {
	x := 43
	y := "Nihad Hossain"
	return x, y

}
