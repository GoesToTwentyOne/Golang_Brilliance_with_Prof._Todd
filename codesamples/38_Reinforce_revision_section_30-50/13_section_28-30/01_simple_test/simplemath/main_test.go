package simplemath

import "testing"

func TestAdd(t *testing.T) {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if x := Add(xi...); x != 45 {
		t.Error("expect 36 got ", x)
	}
}
