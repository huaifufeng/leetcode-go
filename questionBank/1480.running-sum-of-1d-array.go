/**
  题目地址：https://leetcode-cn.com/problems/running-sum-of-1d-array/

  给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
  请返回 nums 的动态和。

  <pre>
  输入：nums = [1,2,3,4]
  输出：[1,3,6,10]
  解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
  </pre>

  <pre>
  输入：nums = [1,1,1,1,1]
  输出：[1,2,3,4,5]
  解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。
  </pre>

  <pre>
  输入：nums = [3,1,2,10,1]
  输出：[3,4,6,16,17]
  </pre>

  PS:
    1 <= nums.length <= 1000
    -10^6 <= nums[i] <= 10^6

  解题： 时间复杂度O(n)
	1.循环数组，逐步计算每个位置的值，记录下来
*/
package questionBank

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}
