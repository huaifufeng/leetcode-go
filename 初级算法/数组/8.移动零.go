//题目地址:https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/28/
package main

import "fmt"

func main() {
	nums := []int{1, 0, 3}
	moveZeroes(nums)

}

func moveZeroes(nums []int)  {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for ; j< len(nums); j++ {
		nums[j] = 0
	}

	fmt.Println(nums)
}