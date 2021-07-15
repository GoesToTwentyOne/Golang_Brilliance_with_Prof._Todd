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
	fmt.Println("---------------------------------------------------------------------------------------")
	delete(m, "bill_kennedy")
	for k, v := range m {
		fmt.Println("Key is ", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}

	fmt.Println("------------------------------------simple search -----------------------------------")

	if v, ok := m["Nihad"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("your query value is not available")
	}
	fmt.Println("------------------------------------simple delete -----------------------------------")
	if v, ok := m["no_dr"]; ok {
		delete(m, `no_dr`)
		fmt.Println("delete value :", v)
	} else {
		fmt.Println("your query for delete value is not available")
	}
}
