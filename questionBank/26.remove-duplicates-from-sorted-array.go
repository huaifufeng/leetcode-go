/**
  题目地址：https://leetcode-cn.com/problems/merge-two-sorted-lists/

  将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

  <pre>
     输入：1->2->4, 1->3->4
     输出：1->1->2->3->4->4
  </pre>

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和pos-1位置的元素的值，如果不相等，就将pos的值替换为当前的元素的值
*/
package questionBank

func removeDuplicates(nums []int) int {
	//空数组直接返回0
	if len(nums) == 0 {
		return 0
	}

	pos := 1
	for i := 1; i < len(nums); i++ {
		//因为是有序的数组，只需要比较当前的元素和替换位置前一个元素的值，如果不相等，说明可以替换到当前位置
		if nums[i] != nums[pos - 1] {
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}