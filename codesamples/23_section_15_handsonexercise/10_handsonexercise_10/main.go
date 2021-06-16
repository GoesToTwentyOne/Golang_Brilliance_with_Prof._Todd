package main

import "fmt"

func main() {
	foo()

}
func foo() {
	x := 45
	for {

		x++
		fmt.Println(x)
		if x == 50 {
			break
		}

	}

}
