package questionBank

import "testing"

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	res := findPeakElement(nums)
	if res != 2 {
		t.Error("寻找峰值算法测试1错误", res)
	}

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	res = findPeakElement(nums)
	if res != 1 {
		t.Error("寻找峰值算法测试2错误", res)
	}

	nums = []int{1}
	res = findPeakElement(nums)
	if res != 0 {
		t.Error("寻找峰值算法测试3错误", res)
	}

	nums = []int{1, 2, 3}
	res = findPeakElement(nums)
	if res != 2 {
		t.Error("寻找峰值算法测试4错误", res)
	}
}
