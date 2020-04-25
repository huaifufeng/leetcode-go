package algorithm

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []string{"abc", "ad", "cdb"}
	RadixSort(nums)

	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("选择排序结果有问题")
		}
	}
}
