package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(x)
	for i, value1 := range x {
		fmt.Println("record : ", i)
		for j, valuefinal := range value1 {
			fmt.Printf(" \t \t index position %v value %v \n", j, valuefinal)

		}

	}

}
