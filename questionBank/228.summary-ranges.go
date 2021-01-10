/**
  题目地址：https://leetcode-cn.com/problems/summary-ranges/

  给定一个无重复元素的有序整数数组 nums 。

  返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，
  并且不存在属于某个范围但不属于 nums 的数字 x 。

  列表中的每个区间范围 [a,b] 应该按如下格式输出：
    "a->b" ，如果 a != b
	"a" ，如果 a == b

  <pre>
     输入：nums = [0,1,2,4,5,7]
	 输出：["0->2","4->5","7"]
	 解释：区间范围是：
	 [0,2] --> "0->2"
	 [4,5] --> "4->5"
	 [7,7] --> "7"
  </pre>

  <pre>
     输入：nums = [0,2,3,4,6,8,9]
	 输出：["0","2->4","6","8->9"]
	 解释：区间范围是：
	 [0,0] --> "0"
	 [2,4] --> "2->4"
	 [6,6] --> "6"
	 [8,9] --> "8->9"
  </pre>

  <pre>
     输入：nums = []
 	 输出：[]
  </pre>

  解题：
    1、因为是有序数组，只需要从前往后依次处理每个元素就可以
*/
package questionBank

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		str := strconv.Itoa(nums[0])
		return []string{str}
	}

	ret := []string{}
	left := nums[0]
	right := nums[0]
	for i := 1; i < len(nums); i++ {
		//大于当前值+1，说明这里有中断，可以组成一个序列
		if nums[i] > right+1 {
			//相等，说明只有单独一个元素
			if left == right {
				str := strconv.Itoa(left)
				ret = append(ret, str)

				left = nums[i]
				right = nums[i]
				continue
			}

			leftStr := strconv.Itoa(left)
			rightStr := strconv.Itoa(right)
			ret = append(ret, leftStr+"->"+rightStr)

			left = nums[i]
			right = nums[i]
			continue
		}

		right = nums[i]
	}

	if left == right {
		str := strconv.Itoa(left)
		ret = append(ret, str)

		return ret
	}

	leftStr := strconv.Itoa(left)
	rightStr := strconv.Itoa(right)
	ret = append(ret, leftStr+"->"+rightStr)

	return ret
}
