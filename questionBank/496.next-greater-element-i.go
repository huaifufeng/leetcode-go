/**
 题目地址：https://leetcode-cn.com/problems/next-greater-element-i/

解题：
  1、循环nums1，依次取出值到nums2中寻找 O(m*n)
  2、

*/
package questionBank

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		cur := 10001
		for _, v2 := range nums2 {
			if v == v2 {
				cur = v
				continue
			}

			if cur < v2 {
				res[i] = v2
				break
			}
		}

		if res[i] == 0 {
			res[i] = -1
		}
	}

	return res
}
