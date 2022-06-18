package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for i := 0; i < 5; i++ {
		fmt.Println(x[i])
	}

}
