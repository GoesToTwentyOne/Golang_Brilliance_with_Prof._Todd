package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	for _, value := range a {
		fmt.Println(value)
	}
}
