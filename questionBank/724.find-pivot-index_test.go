package questionBank

import "testing"

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	if res != 3 {
		t.Error("寻找数组中序错误1！")
	}

	nums = []int{1, 2, 3}
	res = pivotIndex(nums)
	if res != -1 {
		t.Error("寻找数组中序错误2！")
	}

	nums = []int{2, 1, -1}
	res = pivotIndex(nums)
	if res != 0 {
		t.Error("寻找数组中序错误3！")
	}

	nums = []int{0, 0, 0, 0, 1}
	res = pivotIndex(nums)
	if res != 4 {
		t.Error("寻找数组中序错误4！")
	}
}
