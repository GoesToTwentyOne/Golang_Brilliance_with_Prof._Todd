package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	x[`Nadidolapo`] = []string{`Manjaro,Turusko`, `Guava`, `Beach love`}
	//before delete
	fmt.Println("before delete ", x)

	//delete
	delete(x, `moneypenny_miss`)

	for key, rowvalue := range x {
		fmt.Println("record name ", key)
		for i, colvalue := range rowvalue {
			fmt.Printf("\t \t index position %v  value %v\n", i, colvalue)
		}

	}
	fmt.Println("after delete ", x)

}
