package questionBank

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	res := singleNumber(nums)
	if len(res) != 2 {
		t.Error("只出现一次的数字 III算法测试1错误", res)
	}

	nums = []int{-1, 0}
	res = singleNumber(nums)
	if len(res) != 2 {
		t.Error("只出现一次的数字 III算法测试2错误", res)
	}

	nums = []int{0, 1}
	res = singleNumber(nums)
	if len(res) != 2 {
		t.Error("只出现一次的数字 III算法测试2错误", res)
	}
}
