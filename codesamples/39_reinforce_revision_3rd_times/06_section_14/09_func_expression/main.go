package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("function is first class citizen in golang")
	}
	f()
	f2 := func(x int, y int) {
		fmt.Println("function is first class citizen in golang ")
		fmt.Printf("age of golang is %v %v\n", x, y)
	}
	f2(2009, 2022)
}
