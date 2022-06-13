package main

import "fmt"

var x = 20

func main() {
	//format printing
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	//print
	fmt.Print(x)
	fmt.Println()
	//print with new line
	fmt.Println(x)
	//string printing
	s := fmt.Sprint(x)
	//fmt.Sprint(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%v \t %#v \t %x \t %#x \t %b \t ", x, x, x, x, x)

	s = fmt.Sprintf("%v \t %#v \t %x \t %#x \t %b \t ", x, x, x, x, x)
	//fmt.Sprint(s)

}
