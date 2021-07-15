package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["bill_kennedy"] = []string{"Hazelnut", "Grape", "Banana"}

	for k, v := range m {
		fmt.Println("Key is ", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
