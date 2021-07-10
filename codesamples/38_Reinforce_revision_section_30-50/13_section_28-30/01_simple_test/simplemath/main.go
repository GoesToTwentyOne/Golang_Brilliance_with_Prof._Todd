package simplemath

func Add(s ...int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]

	}
	return sum

}
