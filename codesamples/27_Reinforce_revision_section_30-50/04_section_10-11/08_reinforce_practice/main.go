package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	for i, row := range x {
		fmt.Println("Row index : ", i)

		for _, col := range row {
			fmt.Print("\t \t ", col)
		}
		fmt.Println()
	}

}
