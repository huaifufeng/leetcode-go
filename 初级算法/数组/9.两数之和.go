//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/29/
package main

import "fmt"

func main() {
	num := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(num, target)

	fmt.Println(result)
}

func twoSum(num []int, target int)[]int {
	hasMap := make(map[int]int)

	ret := make([]int, 2)

	for i:=0; i<len(num); i++ {
		if _, ok := hasMap[target - num[i]]; ok {
			ret[0] = hasMap[target - num[i]]
			ret[1] = i

			return ret
		} else {
			hasMap[num[i]] = i
		}
	}

	return ret
}