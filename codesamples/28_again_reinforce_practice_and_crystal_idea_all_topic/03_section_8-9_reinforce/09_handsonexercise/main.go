package main

import "fmt"

func main() {
	switch favsport := "Cricket"; favsport {
	case "Football":
		fmt.Println("fav sport Football")
	case "Cricket":
		fmt.Println("fav sport Cricket")
	case "Basket Ball":
		fmt.Println("fav sport Basket Ball")
	default:
		fmt.Println("fav sport is not available")

	}

}
