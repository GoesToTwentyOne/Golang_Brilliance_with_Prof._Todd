package main

import "fmt"

func main() {
	m := map[string]string{
		"Md.Nihad Hossain": "USA",
		"Alex":             "Canada",
		"Zlino":            "Finland",
		"Tyu Niya":         "Russia",
		"Adfin":            "London",
	}
	fmt.Println(m)
	if v, ok := m["Adfin"]; ok {
		fmt.Println(v)
	}
	if _, ok := m["Adfin"]; ok {
		delete(m, "Adfin")
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
