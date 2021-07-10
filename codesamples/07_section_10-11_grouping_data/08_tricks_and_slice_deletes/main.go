package main

import "fmt"

func main() {
	x := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 44, 88, 66}
	fmt.Println(x)
	//tricks delete
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
