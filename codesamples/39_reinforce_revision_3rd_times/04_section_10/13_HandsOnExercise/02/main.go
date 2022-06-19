package main

import "fmt"

func main() {
	a := []int{4, 5, 7, 8, 9, 44, 55, 66, 66, 88}
	for i, v := range a {
		fmt.Printf("index :%v      value: %v\n", i, v)
	}

}
