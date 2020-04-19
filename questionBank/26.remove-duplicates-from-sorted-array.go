/**
  题目地址：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

  给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
  不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
  <pre>
    给定数组 nums = [1,1,2],
    函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
    你不需要考虑数组中超出新长度后面的元素。
  </pre>

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和pos-1位置的元素的值，如果不相等，就将pos的值替换为当前的元素的值
*/
package questionBank

import "fmt"

func removeDuplicates(nums []int) int {
	//空数组直接返回0
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		//因为是有序的数组，只需要比较当前的元素和替换位置前一个元素的值，如果不相等，说明可以替换到当前位置
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	//这里没有计算0位的元素，最后增加上
	i++

	return i
}

//removeDuplicates2 删除数组中重复k次的以上的元素
//这里k需要是大于0的数字，要不0位的数字处理会有问题
func removeDuplicates2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for _, v := range nums {
		//这里i<k 这个条件，会把数组的前k个元素过滤掉不进行处理，最多会重复k次
		//v > nums[i - k] 比较当前元素和之前第k个元素的值，大说明是之前没有的，放到i的下个位置，
		// 如果相等，说明超过重复的次数k，不进行复制，相当于丢弃了
		if i < k || v > nums[i - k] {
			i++
			nums[i] = v
		}

		fmt.Println(i)
	}

	return i
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}