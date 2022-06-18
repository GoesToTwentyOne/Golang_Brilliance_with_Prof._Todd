package main

import "fmt"

func main() {
	x := 43

	if x == 43 {
		fmt.Println("The value is 43")
	} else if x == 32 {
		fmt.Println("The value is 32")
	} else {
		fmt.Println("This is empty")
	}

}
