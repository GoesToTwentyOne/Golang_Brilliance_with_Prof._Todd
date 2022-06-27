package main

import "fmt"

func main() {
	get()
	set("SET")
	r := sets("SETS")
	fmt.Println(r)
	r1, r2 := gets("ONE", "TWO")
	fmt.Println(r1, r2)
}

func get() {
	fmt.Println("Hello I am form GET")
}
func set(s string) {
	fmt.Println("I am from , ", s)
}
func sets(s string) string {
	return fmt.Sprint("I am from ", s)
}
func gets(s, s1 string) (string, bool) {
	a := fmt.Sprint("Hello    ", s, s1)
	b := true
	return a, b
}
