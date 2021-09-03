package interviewQuestion

import (
	"testing"
)

func TestSmallestK(t *testing.T) {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	res := smallestK(arr, 4)
	if len(res) != 4 {
		t.Error("最小K个数算法错误")
	}
}
