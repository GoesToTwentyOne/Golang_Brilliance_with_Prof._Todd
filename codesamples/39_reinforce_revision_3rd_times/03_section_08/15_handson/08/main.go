package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("ture")
		fallthrough
	case 3 == 3:
		fmt.Println("true")
	case 5 == 5:
		fmt.Println("ture")
	case 2 != 2:
		fmt.Println("false")
	case 3 != 3:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}

}
