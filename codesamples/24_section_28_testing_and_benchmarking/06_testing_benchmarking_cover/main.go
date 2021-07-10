package main

import "fmt"

func main() {
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))

}

func FindOutlier(integers []int) int {
	var even int
	var odd int
	var ecount int
	var ocount int

	for i := 0; i < len(integers); i++ {
		if integers[i]%2 == 0 {
			even = integers[i]
			ecount++

		} else {
			odd = integers[i]
			ocount++
		}
	}
	if ecount > 0 {
		return even
	}
	return odd

}
