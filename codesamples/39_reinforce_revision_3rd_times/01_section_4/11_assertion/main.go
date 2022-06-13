package main

import "fmt"

var x interface{} = "Nihad Hossain"

func main() {
	fmt.Printf("%T\n", x)
	s := x.(string)
	fmt.Printf("%v\n", s)
	fmt.Printf("%T\n", s)
	s, ok := x.(string)
	fmt.Printf("%v %v\n", s, ok)
	fmt.Printf("%T\n", s)

}
