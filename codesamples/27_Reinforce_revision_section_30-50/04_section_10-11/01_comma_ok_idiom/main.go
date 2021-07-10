package main

import "fmt"

func main() {
	Teamgone := map[string]int{
		"Nihad":  21,
		"Lisbia": 23,
		"Facina": 19,
	}
	if v, ok := Teamgone["Nasia"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("value dosent exist")
	}

}
