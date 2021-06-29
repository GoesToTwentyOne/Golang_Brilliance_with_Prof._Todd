package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 10000 {
			break

		}
		fmt.Println(x)
		x++

	}
}
