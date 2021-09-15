/**
  题目地址：https://leetcode-cn.com/problems/find-peak-element/

  解题：
  1、最简单的方法，循环数组，依次比较每个元素，O(n)
  2、题目要求时间复杂度O(log n),这里可以使用归并算法，分别计算左右两侧的值，获取其中的一个结果
*/
package questionBank

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	return getIndex(nums, 0, len(nums)-1)
}

func getIndex(nums []int, left, right int) int {
	if right-left == 1 {
		if left == 0 && nums[left] > nums[right] {
			return left
		}

		if right == len(nums)-1 && nums[right] > nums[left] {
			return right
		}

		return -1
	}

	mid := (left + right) / 2
	if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		return mid
	}

	index := getIndex(nums, left, mid)
	if index != -1 {
		return index
	}

	index = getIndex(nums, mid, right)

	return index
}

func findPeakElement1(nums []int) int {
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			return i
		}
	}

	return 0
}
