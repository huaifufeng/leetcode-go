/**
  题目地址：https://leetcode-cn.com/problems/B1IidL/

  解题：
    1、循环这个数组，找到比左右两边大的值，可以循环，也可以二分
*/
package offer

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid-1] {
			if arr[mid] > arr[mid+1] {
				return mid
			} else {
				left = mid
			}
		} else {
			right = mid
		}
	}

	return 0
}
