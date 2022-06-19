package main

import "fmt"

func main() {
	m := map[string]float64{
		"Md.Nihad Hossain": 10.9,
		"Alex":             20.5,
		"Tyne":             54,
	}
	fmt.Println(m)
	m["Samni"] = 444.4
	for k, v := range m {
		fmt.Println(k, v)
	}
}
