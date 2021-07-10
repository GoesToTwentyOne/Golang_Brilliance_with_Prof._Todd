package main

import "fmt"

func main() {
	x := []int{1, 3, 2, 54, 4, 5, 6, 6, 6, 6}
	fmt.Println("length :", len(x))
	fmt.Println("capacity : ", cap(x))
	for i, value := range x {
		fmt.Println(i, " \t value and index \t", value)
	}
}
