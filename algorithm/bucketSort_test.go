package algorithm

import "testing"

func TestBucketSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	BucketSort(nums, 3)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("计数排序结果有问题")
		}
	}
}
