package main

import "fmt"

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
	m["Niad Hossain"] = []string{"Chocklate", "Ice cream", "Bevarage"}

	fmt.Println("---------------------------------------insert after print ---------------------------------------------")
	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t \t ", v)

		}

	}
	if _, ok := m[`bond_james`]; ok {
		delete(m, `bond_james`)
		delete(m, `Chaika`)

	} else {
		fmt.Println("data not foun")
	}
	if _, ok := m[`Chaika`]; ok {

		delete(m, `Chaika`)

	} else {
		fmt.Println("data not found")
	}
	fmt.Println("---------------------------------------delete  after print ---------------------------------------------")
	for key, value := range m {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t \t ", v)

		}

	}

}
