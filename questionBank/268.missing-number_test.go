package questionBank

import "testing"

func TestMissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	res := missingNumber(nums)
	if res != 2 {
		t.Error("丢失的数字算法测试1错误")
	}

	nums = []int{0, 1}
	res = missingNumber(nums)
	if res != 2 {
		t.Error("丢失的数字算法测试2错误")
	}

	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	res = missingNumber(nums)
	if res != 8 {
		t.Error("丢失的数字算法测试3错误")
	}

	nums = []int{0}
	res = missingNumber(nums)
	if res != 1 {
		t.Error("丢失的数字算法测试4错误")
	}
}
