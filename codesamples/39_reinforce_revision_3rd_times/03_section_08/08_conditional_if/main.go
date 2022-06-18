package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is Nihad Hossain")
	}
	if false {
		fmt.Println("Working with China")
	}
	if !true {
		fmt.Println("Working with Army")
	}
	if !false {
		fmt.Println("Working with friends")
	}
	if !(!true) {
		fmt.Println("Working with Alex")
	}
	if 2 == 2 {
		fmt.Println("this is 2")
	}
	if 2 != 2 {
		fmt.Println("This is not equal to 2")
	}
	if x := 32; x == 32 {
		fmt.Println("This is 32")
	}
}
