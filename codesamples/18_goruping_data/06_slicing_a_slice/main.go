package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x[:])
	fmt.Println(x[1:4])
	fmt.Println(x[:3])
	fmt.Println("--------------------------------------------------------------------------------------------- ")
	for i := 0; i < len(x); i++ {
		fmt.Println(i, "------------------", x[i])
	}
}
