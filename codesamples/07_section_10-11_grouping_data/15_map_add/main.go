package main

import "fmt"

func main() {
	x := map[string]int{
		"Nihad":  21,
		"Rowja":  7,
		"Tpanro": 21,
		"Alex":   25,
	}
	x["Jloneya"] = 52
	fmt.Println(x)
	for key, value := range x {
		fmt.Println(key, value)
	}

}
