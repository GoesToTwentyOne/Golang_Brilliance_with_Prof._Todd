package main

import "fmt"

func main() {
	x := [2][3]string{{"Niahd", "Shova", "Tanota"}, {"Jerry", "Scott", "Pansalvia"}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("x[%d][%d] = %s \n", i, j, x[i][j])
		}
	}
}
