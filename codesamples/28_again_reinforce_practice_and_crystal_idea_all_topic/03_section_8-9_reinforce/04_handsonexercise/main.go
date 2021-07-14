package main

import "fmt"

func main() {
	x := 1999
	count := 0
	for {
		if x > 2021 {
			break
		}
		count++
		fmt.Println(x)
		x++

	}
	fmt.Println("Total alive ", count)
}
