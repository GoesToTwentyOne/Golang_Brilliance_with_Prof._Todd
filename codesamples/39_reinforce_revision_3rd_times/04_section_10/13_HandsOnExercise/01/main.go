package main

import "fmt"

func main() {
	a := [5]int{4, 5, 7, 8, 9}
	for i, v := range a {
		fmt.Printf("index :%v      value: %v\n", i, v)
	}

}
