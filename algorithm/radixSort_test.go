package algorithm

import (
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []string{"bda", "cfd", "qwe", "yui", "abc", "rrr", "uue"}
	RadixSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("选择排序结果有问题")
		}
	}
}
