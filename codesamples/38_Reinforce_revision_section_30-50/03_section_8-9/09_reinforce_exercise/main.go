package main

import "fmt"

func main() {
	switch favsport := "cricket"; favsport {
	case "football":
		fmt.Println("favsport: football")
	case "racket":
		fmt.Println("favsport : Racket")
	case "serfing":
		fmt.Println("favsport: serfing")
	case "cricket":
		fmt.Println("favsport : cricket")
	default:
		fmt.Println("done")

	}
}
