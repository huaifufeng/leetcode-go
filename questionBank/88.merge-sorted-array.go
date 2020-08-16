/**
  题目地址：https://leetcode-cn.com/problems/merge-sorted-array/

  给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
  说明：
    初始化nums1 和 nums2 的元素数量分别为m 和 n。
    你可以假设nums1有足够的空间（空间大小大于或等于m + n）来保存 nums2 中的元素。


  <pre>
    输入:
      nums1 = [1,2,3,0,0,0], m = 3
      nums2 = [2,5,6],       n = 3

    输出:[1,2,2,3,5,6]
  </pre>

  解题：
	1、这里nums1已经包含了两个数组的长度，只需要从后向前，按照顺序把两个数组的元素放到nums1中
*/
package questionBank

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	cur := len(nums1) - 1
	//只有都大于等于0时才进行处理
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}

		cur--
	}

	//如果nums2没有处理完毕，就继续把nums2放入到nums1的相应位置
	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}
}
