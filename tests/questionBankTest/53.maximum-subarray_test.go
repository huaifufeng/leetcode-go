package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	res := questionBank.MaxSubArray(nums)
	if res != 6 {
		t.Error("最大子串和计算错误")
	}
}