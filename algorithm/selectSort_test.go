package algorithm

import "testing"

func TestSelectSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	SelectSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("选择排序结果有问题")
		}
	}
}
