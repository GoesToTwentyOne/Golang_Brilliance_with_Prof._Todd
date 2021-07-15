package main

import "fmt"

func main() {
	xi := []int{4, 6, 7, 8, 9, 9, 9}
	fmt.Println(xi)
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Printf("%T \n", xi)
}
