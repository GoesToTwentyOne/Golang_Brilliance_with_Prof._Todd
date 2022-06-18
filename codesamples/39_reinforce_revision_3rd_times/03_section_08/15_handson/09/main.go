package main

import "fmt"

func main() {
	favSports := "Cricket"
	switch favSports {
	case "Footbal":
		fmt.Println("Football")
	case "Cricket":
		fmt.Println("Cricket")
	case "Hocky":
		fmt.Println("Hocky")
	default:
		fmt.Println("Default")
	}

}
