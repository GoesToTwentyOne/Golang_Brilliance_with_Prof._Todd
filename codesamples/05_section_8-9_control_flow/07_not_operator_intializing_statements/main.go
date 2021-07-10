package main

import "fmt"

func main() {
	if true {
		fmt.Println("one")
	}
	if false {
		fmt.Println("Two")
	}
	if z := 5; z == 6 { //init
		fmt.Println("three")
	}
	if !(2 == 2) {
		fmt.Println("four")
	}
	if !(2 != 2) {
		fmt.Println("five")
	}

}
