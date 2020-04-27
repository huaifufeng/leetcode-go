package algorithm

import "testing"

func TestCockTailSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	CockTailSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("鸡尾酒排序结果有问题")
		}
	}
}

func TestCockTailSort2(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	CockTailSort2(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("鸡尾酒排序2结果有问题")
		}
	}
}
