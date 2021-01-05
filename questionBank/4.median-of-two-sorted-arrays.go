/**
  题目地址：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

  给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
  进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

  <pre>
     输入：nums1 = [1,3], nums2 = [2]
	 输出：2.00000
	 解释：合并数组 = [1,2,3] ，中位数 2
  </pre>

  注意：
    假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

  解题：
    1、使用框架自带的函数，合并两个切片之后排序，然后获取中位数的值
*/
package questionBank

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)

	sort.Ints(nums)

	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}
