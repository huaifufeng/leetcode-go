package interviewQuestion

import (
	"testing"
)

func TestSortedMergeLcci(t *testing.T) {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3

	ret := Merge(A, m, B, n)

	if ret[5] != 6 {
		t.Error("结果不对")
	}
}
