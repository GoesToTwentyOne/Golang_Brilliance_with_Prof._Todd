package mysum

func Sum(n ...int) int {
	s := 0
	for _, value := range n {
		s += value
	}
	return s

}
