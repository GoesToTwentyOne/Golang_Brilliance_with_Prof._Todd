package main

import "fmt"

func main() {
	switch x := 5; x {
	case 1, 2, 3:
		fmt.Println("this is 1 or 2 or 3")

	case 4, 5, 6:
		fmt.Println("this is 4 or 5 or 6")
		fallthrough

	case 7, 8, 9:
		fmt.Println("this is 7 or 8 or 9")
		fallthrough
	case 10, 11, 13:
		fmt.Println("this is 7 or 8 or 9")
		fallthrough
	default:
		fmt.Println("deafault value")
	}
}
