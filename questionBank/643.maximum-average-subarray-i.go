/**
  题目地址：https://leetcode-cn.com/problems/maximum-average-subarray-i/

  给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。


  <pre>
     输入：[1,12,-5,-6,50,3], k = 4
     输出：12.75
     解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
  </pre>

  注意：
    1 <= k <= n <= 30,000。
    所给数据范围 [-10,000，10,000]。

  解题： 时间复杂度

*/
package questionBank

func findMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	avg := sum / float64(k)
	for j := k; j < len(nums); j++ {
		sum -= float64(nums[j-k])
		sum += float64(nums[j])

		temp := sum / float64(k)

		if temp > avg {
			avg = temp
		}
	}

	return avg
}
