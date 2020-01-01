/**
  题目地址：https://leetcode-cn.com/problems/search-insert-position/

  给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
  如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
  你可以假设数组中无重复元素。

  <pre>
    输入: [1,3,5,6], 5
    输出: 2
  </pre>

  <pre>
    输入: [1,3,5,6], 7
    输出: 4
  </pre>

  解题：
    这里就是在数组中查找等于target或第一个大于target元素的位置
    1、循环遍历整个数组，查找指定的位置 O(n)
    2、二分查找指定的元素， O(logn)
*/
package questionBank

func searchInsert(nums []int, target int) int {
	length := len(nums)
	//对特殊的情况进行处理
	//1、如果目标比数组的第一个元素都小，直接返回0
	//2、如果目标比最后一个元素的值都大，直接返回数组长度
	if target <= nums[0] {
		return 0
	} else if target > nums[length - 1]{
		return length
	}

	left := 0
	right := length - 1

	//只有在左值小于右值的情况下进行处理，
	for left < right  {
		//这里是左中值
		mid := left + (right - left) / 2
		//相等就是找到目标元素位置，直接返回
		if nums[mid] == target {
			return mid
		}

		//如果中值比目标小，说明要找的位置在右边，增加左值的值
		//如果比目标大，说明在左边或者就是中值的位置，所以把这个位置赋值给right
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}