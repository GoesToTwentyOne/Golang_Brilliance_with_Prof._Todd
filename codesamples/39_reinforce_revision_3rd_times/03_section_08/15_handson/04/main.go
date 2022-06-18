package main

import "fmt"

func main() {
	x := 21

	for {
		if x > 65 {
			break
		}
		fmt.Println(x)
		x++
	}

}
