package main

import "fmt"

func main() {
	z := 1
	for {
		if z > 9 {
			break
		}
		fmt.Printf("%v\n", z)
		z++
	}
}
