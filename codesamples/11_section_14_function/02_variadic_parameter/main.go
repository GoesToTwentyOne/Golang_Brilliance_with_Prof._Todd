package main

import "fmt"

func main() {
	goes(1, 2, 3, 4, 4, 56, 678, 88777, 7777665532, 223)

}
func goes(n ...int) {
	// fmt.Println(n)
	sum := 0

	for i, value := range n {

		sum += value
		fmt.Println(i, " index of value ", value, "sum of order ", sum)

	}

}
