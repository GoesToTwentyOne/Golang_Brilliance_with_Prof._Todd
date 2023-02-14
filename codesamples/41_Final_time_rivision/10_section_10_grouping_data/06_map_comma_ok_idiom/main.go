package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Md. Nihad Hossian": 21,
		"Alex Goot":         21,
		"Gertoy":            23,
		"Uniktri":           24,
	}
	//v, ok := x["Alex Goot"]
	//fmt.Println(v)
	//fmt.Println(ok)
	//fmt.Println(x)
	x["Otolia"] = 24
	if v, ok := x["Uniktri"]; ok {
		fmt.Println(v)
	}
	for v, k := range x {
		fmt.Println(v, k)
	}
	if v, ok := x["Uniktri"]; ok {
		fmt.Println(v)
		delete(x, "Uniktri")
	}
	for v, k := range x {
		fmt.Println(v, k)
	}
}
