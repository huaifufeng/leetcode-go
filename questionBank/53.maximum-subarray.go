/**
  题目地址：https://leetcode-cn.com/problems/maximum-subarray/

  给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

  <pre>
     输入: [-2,1,-3,4,-1,2,1,-5,4],
     输出: 6
     解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
  </pre>

  解题：
	1、依次循环整个数组，当之前元素加和之后的结果是负数，说明会对和产生负效果，就舍弃这部分计算，以当前元素开始进行新的计算
       同时，对加和的结果和保存的最大分组和进行比较，保留其中最大的部分    时间复杂度O(n)
*/
package questionBank

func maxSubArray(nums []int) int {
	maxValue := nums[0]
	sum := 0
	for _, val := range nums {
		if sum > 0 {
			sum += val
		} else {
			sum = val
		}

		//将之前保存的最大值与加和结进行比较，取其中的最大值
		if maxValue < sum {
			maxValue = sum
		}
	}

	return maxValue
}
