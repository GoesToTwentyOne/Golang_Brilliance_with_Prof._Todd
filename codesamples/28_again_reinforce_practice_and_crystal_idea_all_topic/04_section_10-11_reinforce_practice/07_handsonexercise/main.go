package main

import "fmt"

func main() {
	xxs := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(xxs)
	for i, row := range xxs {
		fmt.Println("Rows index :", i)
		for i, col := range row {
			fmt.Println("\t", i, col)
		}

	}
}
