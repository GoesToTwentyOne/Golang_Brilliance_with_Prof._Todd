package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	defer foo(xi...)
	bar(xi)

}
func foo(n ...int) {
	sum := 0
	for _, value := range n {
		sum += value
	}
	fmt.Println("foo sum : ", sum)

}
func bar(n []int) {
	sum := 0
	for _, value := range n {
		sum += value
	}
	fmt.Println("bar sum : ", sum)

}
