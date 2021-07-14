package main

import "fmt"

func main() {
	x := 10
	for {
		if x > 100 {
			break
		}
		// if x%4 == 0 {
		// 	fmt.Println("divisible by 4 :", x)
		// }
		fmt.Printf("%d is divide and modulaus %d \n", x, x%4)
		x++
	}
}
