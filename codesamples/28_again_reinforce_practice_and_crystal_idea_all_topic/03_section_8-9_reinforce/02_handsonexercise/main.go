package main

import "fmt"

func main() {
	x := 65
	for {
		if x > 90 {
			break
		}
		fmt.Println("Ascii decimal ", x)
		for i := 1; i <= 3; i++ {
			fmt.Printf(" \t \t %#U \n", x)
		}
		x++
	}
}
