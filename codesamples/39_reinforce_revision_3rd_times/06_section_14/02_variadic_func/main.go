package main

import "fmt"

func main() {
	goes(3, 4, 5, 67, 8, 6, 7, 8)

}
func goes(x ...int) {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("position is %d  sum is %d\n", i, sum)

	}

}
