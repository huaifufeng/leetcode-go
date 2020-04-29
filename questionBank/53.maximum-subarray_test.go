package questionBank

import (
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	if res != 6 {
		t.Error("最大子串和计算错误")
	}
}
