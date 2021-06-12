package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// fmt.Println(x)
	for key, rowvalue := range x {
		fmt.Println("Recoded row value ", key)
		for i, colvalue := range rowvalue {
			fmt.Printf("\t \t index position %v value %v \n", i, colvalue)

		}
	}
}
