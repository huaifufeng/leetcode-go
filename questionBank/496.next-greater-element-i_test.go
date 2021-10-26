package questionBank

import "testing"

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	res := nextGreaterElement(nums1, nums2)
	if res[0] != -1 || res[1] != 3 || res[2] != -1 {
		t.Error("下一个更大元素 I算法测试1错误", res)
	}

	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}

	res = nextGreaterElement(nums1, nums2)
	if res[0] != 3 || res[1] != -1 {
		t.Error("下一个更大元素 I算法测试2错误", res)
	}
}
