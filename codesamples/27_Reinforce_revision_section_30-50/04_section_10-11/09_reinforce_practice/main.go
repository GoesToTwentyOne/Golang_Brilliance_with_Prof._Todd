package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t \t ", v)

		}

	}
}
