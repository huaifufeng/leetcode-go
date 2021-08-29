/**
  题目地址：https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/

  给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。
  子数组 定义为原数组中的一个连续子序列。
  请你返回 arr 中 所有奇数长度子数组的和 。

  <pre>
    输入：arr = [1,4,2,5,3]
    输出：58
    解释：所有奇数长度子数组和它们的和为：
    [1] = 1
    [4] = 4
    [2] = 2
    [5] = 5
    [3] = 3
    [1,4,2] = 7
    [4,2,5] = 11
    [2,5,3] = 10
    [1,4,2,5,3] = 15
    我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
  </pre>

  <pre>
    输入：arr = [1,2]
    输出：3
    解释：总共只有 2 个长度为奇数的子数组，[1] 和 [2]。它们的和为 3 。
  </pre>

  <pre>
    输入：arr = [10,11,12]
    输出：66
  </pre>

  PS:
    1 <= arr.length <= 100
    1 <= arr[i] <= 1000
  解题：
  1、暴力循环数组，O(n^3)
*/
package questionBank

func sumOddLengthSubarrays(arr []int) int {
	maxOdd := len(arr)
	if len(arr)%2 == 0 {
		maxOdd = len(arr) - 1
	}

	sum := 0
	for index, _ := range arr {
		if index == len(arr)-1 {
			sum += arr[index]
			break
		}
		for i := 1; i <= maxOdd; i += 2 {
			if index+i > cap(arr) {
				break
			}
			subArr := arr[index : index+i]
			for _, subValue := range subArr {
				sum += subValue
			}
		}
	}

	return sum
}
