/**
  题目地址：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

  给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
  如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
  注意你不能在买入股票前卖出股票。

  <pre>
    输入: [7,1,5,3,6,4]
    输出: 5
    解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
  </pre>

  解题：
	1、第一种就是暴力破解，依次对比数组的每个元素和之后的元素 O(n^2)
*/
package questionBank

func maxProfit(prices []int) int {
	maxPrice := 0

	//对所有数组元素进行比较
	for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				price := prices[j] - prices[i]
				if maxPrice < price {
					maxPrice = price
				}
			}
		}
	}

	return maxPrice
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}