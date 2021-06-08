package main

import "fmt"

func main() {
	a := (42 == 41)
	b := (42 <= 41)
	c := (42 >= 41)
	d := (42 != 41)
	e := (42 < 43)
	f := (46 > 43)
	fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v", a, b, c, d, e, f)
}
