/**
  题目地址：https://leetcode-cn.com/problems/find-pivot-index/

  给你一个整数数组nums，请编写一个能够返回数组 “中心索引” 的方法。
  数组 中心索引 是数组的一个索引，其左侧所有元素相加的和等于右侧所有元素相加的和。
  如果数组不存在中心索引，返回 -1 。如果数组有多个中心索引，应该返回最靠近左边的那一个。
  注意：中心索引可能出现在数组的两端。

  解题： 时间复杂度O(n)
	1.推理条件成成立的条件：2 * leftSum + curr = total
*/
package questionBank

func pivotIndex(nums []int) int {
	total := 0
	for _, i2 := range nums {
		total += i2
	}

	leftSum := 0
	for i, val := range nums {
		if 2*leftSum+val == total {
			return i
		}

		leftSum += val
	}

	return -1
}
