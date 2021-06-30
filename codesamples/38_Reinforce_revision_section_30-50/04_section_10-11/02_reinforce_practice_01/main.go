package main

import "fmt"

func main() {
	x := [5]int{}
	for i := 0; i < len(x); i++ {
		x[i] = i + 2
	}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
