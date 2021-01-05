package questionBank

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	if res != 2.0 {
		t.Error("寻找有序数组中位数错误1")
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 2.5 {
		t.Error("寻找有序数组中位数错误2")
	}

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 0.0 {
		t.Error("寻找有序数组中位数错误3")
	}

	nums1 = []int{}
	nums2 = []int{1}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 1.0 {
		t.Error("寻找有序数组中位数错误4")
	}

	nums1 = []int{2}
	nums2 = []int{}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 2.0 {
		t.Error("寻找有序数组中位数错误5")
	}
}
