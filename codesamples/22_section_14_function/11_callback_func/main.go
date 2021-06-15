package main

import "fmt"

func main() {
	nslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(nslice...)
	fmt.Println("total n sum", s1)
	s2 := callbackEven(sum, nslice...)
	fmt.Println("even sum ", s2)
	s3 := callbackOdd(sum, nslice...)
	fmt.Println("odd sum ", s3)

}
func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}
func callbackEven(f func(n ...int) int, y ...int) int {
	var yi []int
	for _, value := range y {
		if value%2 == 0 {
			yi = append(yi, value)
		}

	}

	return f(yi...)
}
func callbackOdd(f func(n ...int) int, z ...int) int {
	var zi []int
	for _, value := range z {
		if value%2 != 0 {
			zi = append(zi, value)
		}

	}

	return f(zi...)

}
