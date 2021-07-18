package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sum all ", foo(xi...))
	fmt.Println("even sum ", bar(xi))

}
func foo(n ...int) int {
	evsum := 0
	for _, v := range n {

		evsum += v

	}
	return evsum

}
func bar(n []int) int {
	evsum := 0
	for _, v := range n {
		if v%2 == 0 {
			evsum += v
		}
	}
	return evsum

}
