//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/22/
package 数组

func maxProfit(prices []int) int {
	price := 0
	if len(prices) < 2 {
		return price
	}

	for i:=1;i<len(prices);i++ {
		if prices[i] > prices[i-1] {
			price += prices[i] - prices[i-1]
		}
	}

	return price
}
