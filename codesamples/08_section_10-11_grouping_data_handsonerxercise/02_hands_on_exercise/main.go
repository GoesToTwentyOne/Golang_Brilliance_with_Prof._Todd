package main

import "fmt"

func main() {
	slicein := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range slicein {
		fmt.Println(i, "----- ", value)
	}
	fmt.Printf("%T", slicein)

}
