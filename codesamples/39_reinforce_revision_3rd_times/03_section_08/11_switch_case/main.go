package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2, 3, 4:
		fmt.Println("It was 1,2,3,4 Teir")
	case 5, 6, 7, 8:
		fmt.Println("It was 5,6,7,8 Teir")
	case 9, 10, 11, 12:
		fmt.Println("It was 9,10,11,12 Teir")
	default:
		fmt.Println("Default")

	}

}
