package main

import "fmt"

func main() {
	x := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		x[i] = fmt.Sprintf("Position %d \t", i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println("------------------------------------1st--------------------------------------------")
	fmt.Println(cap(x))
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`,
		` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println("-------------------------------------2nd-------------------------------------------")
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {

		fmt.Println(i, x[i])
	}
}
