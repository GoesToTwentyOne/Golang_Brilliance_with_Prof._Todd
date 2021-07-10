package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("not executed")

	case true:
		fmt.Println("executed")

	default:
		fmt.Println("Done")
	}
}
