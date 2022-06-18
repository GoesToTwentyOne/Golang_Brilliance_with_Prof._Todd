package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x[:3])
	fmt.Println(x[3:])
	fmt.Println(x)
	fmt.Println(x[1:4])
}
