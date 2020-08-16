package offer

import "testing"

func TestMissingNumber(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	res := missingNumber(nums)

	if res != 8 {
		t.Error("寻找缺失元素判断错误")
	}

}
