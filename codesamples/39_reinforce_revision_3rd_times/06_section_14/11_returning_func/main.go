package main

import "fmt"

func main() {
	fmt.Println(goes(21)())

}

func goes(x int) func() int {
	return func() int {
		return x
	}
}
