package questionBank

import (
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{0, 2, 3}
	merge(nums1, len(nums1), []int{1}, 1)

	for k, v := range nums1 {
		if v != k+1 {
			t.Error("合并有序数组方法失败")
		}
	}
}
