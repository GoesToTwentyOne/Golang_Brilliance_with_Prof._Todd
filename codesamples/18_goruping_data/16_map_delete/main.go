package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Nihad":  21,
		"Rowja":  7,
		"Tpanro": 21,
		"Alex":   25,
	}
	x["Jloneya"] = 52
	fmt.Println(x)
	v, ok := x["trafer"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := x["Alex"]; ok {
		fmt.Println(v)
		delete(x, "Alex")
	}
	fmt.Println(x)

}
