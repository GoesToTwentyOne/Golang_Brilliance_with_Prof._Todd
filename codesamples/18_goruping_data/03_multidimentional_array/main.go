package main

import "fmt"

func main() {
	x := [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	for _, row := range x {
		for _, col := range row {
			fmt.Println(col)
		}
	}
}
