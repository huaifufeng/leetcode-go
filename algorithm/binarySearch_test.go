package algorithm

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	ret1 := BinarySearch(nums, 3)

	if ret1 != 2 {
		t.Error("二分查找获取位置错误")
	}

	ret2 := BinarySearch(nums, 8)
	if ret2 != -1 {
		t.Error("二分查找获取位置错误")
	}
}

func TestBinarySearchFirst(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 4, 5, 6}
	ret := BinarySearchFirst(nums, 3)

	if ret != 2 {
		t.Error("二分查找第一次出现获取位置错误")
	}
}

func TestBinarySearchLast(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 4, 5, 6}
	ret := BinarySearchLast(nums, 3)

	if ret != 4 {
		t.Error("二分查找第一次出现获取位置错误")
	}
}

func TestBinarySearchFirstLarge(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5, 6, 7}
	ret := BinarySearchFirstLarge(nums, 5)

	if ret != 5 {
		t.Error("二分查找第一次出现获取位置错误")
	}
}

func TestBinarySearchLastMin(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 6, 7}
	ret := BinarySearchLastMin(nums, 5)

	if ret != 4 {
		t.Error("二分查找第一次出现获取位置错误")
	}
}
