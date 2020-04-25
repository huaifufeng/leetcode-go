package algorithm

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	InsertSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("插入排序结果有问题")
		}
	}
}

func TestInsertSort2(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	InsertSort2(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("插入排序2结果有问题")
		}
	}
}
