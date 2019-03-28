//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/22/
package main

import "fmt"

func main() {
	prices := []int{7,1,5,3,6,4}
	result := maxProfit(prices)

	fmt.Println(result)
}

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
