package main

import "fmt"

func main() {
	switch x := "Nihad"; x {
	case "Alex":
		fmt.Println("Alex")
	case "Nihad":
		fmt.Println("Nihad")
	case "Jony":
		fmt.Println("Jony")
	default:
		fmt.Println("Defult")
	}
}
