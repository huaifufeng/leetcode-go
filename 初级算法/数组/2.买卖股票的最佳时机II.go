package main

import "fmt"

func main() {
	prices := []int{1,2,3,4,5}
	result := maxProfit(prices)

	fmt.Println(result)
}

func maxProfit(prices []int) int {
	price := 0
	if len(prices) < 2 {
		return price
	}

	var listNums []int
	for i:=0;i<len(prices);i++ {
		if len(listNums) == 0 {
			listNums = append(listNums, prices[i])
		} else if len(listNums) == 1 {
			if listNums[0] < prices[i] {
				listNums = append(listNums, prices[i])
			} else {
				listNums[0] = prices[i]
			}
		} else {
			if listNums[1] < prices[i] {
				listNums[1] = prices[i]
			} else {
				price = price + listNums[1] - listNums[0]
				listNums = listNums[0:1]
				listNums[0] = prices[i]
			}
		}
	}

	if len(listNums) == 2 {
		price += listNums[1] - listNums[0]
	}

	return price
}
