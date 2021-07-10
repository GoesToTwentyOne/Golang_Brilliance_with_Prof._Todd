package main

import "fmt"

func main() {
	a := 100
	for {
		a--
		if a == 1 {
			break
		}
		if a%2 != 0 {
			continue

		}
		fmt.Println(a)

	}

}
