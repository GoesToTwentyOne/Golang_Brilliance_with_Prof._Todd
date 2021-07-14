package main

import "fmt"

func main() {
	switch {
	case (2 == 1):
		fmt.Println("is there (2==1)     true or false |answers is ", (2 == 1))
	case (4 == 5):
		fmt.Println("is there (4==5)     true or false |answers is ", (4 == 5))
	// case (4 == 4):
	// 	fmt.Println("is there (4==4)     true or false |answers is ", (4 == 4))
	default:
		fmt.Println("everything is fair and lovely")
	}

}
