package questionBank

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	res := search(nums, target)
	if res != 4 {
		t.Error("二分查找算法测试1错误！")
	}

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 2
	res = search(nums, target)
	if res != -1 {
		t.Error("二分查找算法测试2错误！")
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
