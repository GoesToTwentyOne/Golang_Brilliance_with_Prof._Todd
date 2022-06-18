package main

import "fmt"

func main() {
	var i, j int
	x := [][]string{
		{"Name", "Qualification", "University"},
		{"Md. Nihad Hossain", "B.Sc in Engg", "Bangladesh Army University of Science and Technology"},
		{"Neha", "B.Sc in Engg", "Bangladesh Army University of Science and Technology"},
		{"Alex", "B.Sc in Engg", "Bangladesh Army University of Science and Technology"},
	}

	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {
			fmt.Printf(" %v ", x[i][j])
		}
		fmt.Println()

	}
	fmt.Println(x)
}
