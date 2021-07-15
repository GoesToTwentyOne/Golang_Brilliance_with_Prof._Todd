package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
