/**
  题目地址：https://leetcode-cn.com/problems/teemo-attacking/


  解题：
	1、循环计算每个时间间隔，单独计算长度为1的情况。最后一个元素的中毒时间为间隔时间。O(n)
*/
package questionBank

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 1 {
		return duration
	}

	length := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		end := timeSeries[i] + duration - 1
		if end >= timeSeries[i+1] {
			length += timeSeries[i+1] - timeSeries[i]
		} else {
			length += duration
		}
	}

	length += duration

	return length
}
