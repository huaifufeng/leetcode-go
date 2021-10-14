package offer

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	arr := []int{0, 1, 0}
	res := peakIndexInMountainArray(arr)
	if res != 1 {
		t.Error("山峰数组的顶部算法测试1错误")
	}

	arr = []int{1, 3, 5, 4, 2}
	res = peakIndexInMountainArray(arr)
	if res != 2 {
		t.Error("山峰数组的顶部算法测试2错误")
	}

	arr = []int{0, 10, 5, 2}
	res = peakIndexInMountainArray(arr)
	if res != 1 {
		t.Error("山峰数组的顶部算法测试3错误")
	}

	arr = []int{3, 4, 5, 1}
	res = peakIndexInMountainArray(arr)
	if res != 2 {
		t.Error("山峰数组的顶部算法测试4错误")
	}

	arr = []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	res = peakIndexInMountainArray(arr)
	if res != 2 {
		t.Error("山峰数组的顶部算法测试5错误")
	}
}
