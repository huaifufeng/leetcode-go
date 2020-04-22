package algorithm

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{1,3,2,7,3,5,1,4}
	BubbleSort(nums)

	for i:=0; i< len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("排序结果有问题")
		}
	}
}
