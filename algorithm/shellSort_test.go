package algorithm

import "testing"

func TestShellSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	ShellSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("希尔排序结果有问题")
		}
	}
}

func TestShellSort2(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	ShellSort2(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("希尔排序2结果有问题")
		}
	}
}
