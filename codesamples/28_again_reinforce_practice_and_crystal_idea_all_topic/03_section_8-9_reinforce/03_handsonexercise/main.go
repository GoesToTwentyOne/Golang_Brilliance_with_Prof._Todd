package main

import "fmt"

func main() {
	x := 1999
	count := 0
	for x <= 2021 {
		// if true {
		// 	count++
		// }
		count++

		fmt.Println(x)
		x++

	}
	fmt.Println("total alive ", count)

}
