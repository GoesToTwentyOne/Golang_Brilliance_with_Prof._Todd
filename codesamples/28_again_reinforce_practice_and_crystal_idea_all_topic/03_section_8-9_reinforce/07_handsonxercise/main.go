package main

import "fmt"

func main() {
	x := 35
	for {
		if x < 0 {
			break
		}
		if x == 0 {
			break
		} else {
			fmt.Println(x)

		}

		x--
	}
}
