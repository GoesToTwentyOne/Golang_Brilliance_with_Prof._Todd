package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(x)
	for i, row := range x {
		fmt.Println("record : ", i)
		for j, col := range row {
			fmt.Printf(" \t \t index position %v value %v \n", j, col)

		}

	}

}
