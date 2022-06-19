package main

import "fmt"

func main() {
	m := map[string]int{
		"Md.Nihad Hossain": 23,
		"Alex Goot":        22,
		"Shakira":          38,
		"Rish Mantu":       55,
	}
	fmt.Println(m)
	fmt.Println(m["Shakira"])
	fmt.Printf("\n\n\n\n\n\n\n")
	//invalid printing
	fmt.Println(m["Antic"])
	fmt.Printf("\n\n\n\n\n\n\n")
	//coma ok idiom
	// v, ok := m["Antic"]
	// fmt.Println(ok)
	// fmt.Println(v)
	if v, ok := m["Antic"]; ok {
		fmt.Println(v)
	}

	fmt.Printf("\n\n\n\n\n\n\n")
	if v, ok := m["Alex Goot"]; ok {
		fmt.Println(v)
	}

}
