package main

import "fmt"

func main() {
	x := [][]string{
		{"Go lover", "Go exploder", "Go coder", "Go Learner"},
		{"Nihad Hossain1", "Nihad Hossai2", "Nihad Hossain3 ", "Nihad Hossain4"},
		{"Prolem solving URI", "Prolem solving Codefoce", "Prolem solving Codewars", "Prolem codecademy"},
	}
	for row, value := range x {
		fmt.Println("\t \t", row)
		for col, v := range value {
			fmt.Println(col, v)
		}

	}
}
