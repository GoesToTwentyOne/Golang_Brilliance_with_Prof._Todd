package main

import "fmt"

func main() {
	f := func(n int) int {
		fact := 1
		for i := n; i > 0; i-- {
			fact *= i

		}
		return fact

	}(6)
	fmt.Println(f)

}
