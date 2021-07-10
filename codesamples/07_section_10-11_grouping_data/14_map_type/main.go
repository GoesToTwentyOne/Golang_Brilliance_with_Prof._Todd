package main

import "fmt"

func main() {
	x := map[string]int{
		"Nihad":  21,
		"Rowja":  7,
		"Tpanro": 21,
		"Alex":   25,
	}
	fmt.Println(x)
	fmt.Println(x["Nihad"])
	fmt.Println("Not assign", x["l"])
	v, ok := x["l"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := x["l"]; ok {
		fmt.Println("check then run ", v)
	}

}
