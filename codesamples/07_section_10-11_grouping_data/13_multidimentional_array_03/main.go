package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := [3][10]int{}
	for row := 0; row < 3; row++ {
		for col := 0; col < 10; col++ {
			x[row][col] = rand.Intn(1000000)

		}
	}
	for row := 0; row < 3; row++ {
		fmt.Println()

		for col := 0; col < 10; col++ {
			fmt.Printf("x[%d][%d] = %d ", row, col, x[row][col])

		}
	}

}
