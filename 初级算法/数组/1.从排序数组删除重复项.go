//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/21/
package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,2,3,4,4,4,5,5,6}

	result := removeDuplicates(nums)

	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	index := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[index] = nums[i]
			continue
		}

		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}