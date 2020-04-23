package algorithm

import "testing"

func TestMergeSort(t *testing.T) {
	nums := []int{1,3,2,7,3,5,1,4}
	MergeSort(nums, 0, len(nums) - 1)

	for i:=0; i< len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("排序结果有问题")
		}
	}
}