package main

import "fmt"

func main() {
	x := make([]int, 12, 18)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i, value := range x {

		fmt.Println(value)
		x[i] = i
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 1; i < 4; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
