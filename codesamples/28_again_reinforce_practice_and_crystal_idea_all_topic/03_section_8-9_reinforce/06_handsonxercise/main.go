package main

import "fmt"

func main() {
	x := 35
	for {
		if x < 0 {
			break
		}
		fmt.Println(x)
		x--
	}
}
