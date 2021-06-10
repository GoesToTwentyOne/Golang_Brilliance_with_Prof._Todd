package main

import "fmt"

func main() {
	switch favSport := "football"; favSport {
	case "volyball":
		fmt.Println("this is volyball")
	case "football":
		fmt.Println("this is football")

	case "cricket":
		fmt.Println("this is cricket")
	default:

		fmt.Println("done")

	}

}
