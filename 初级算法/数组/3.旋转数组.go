//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/23/
package main

import "fmt"

func main(){
	array1 := []int{1,2,3}
	array2 := []int{-1,-100,3,99}
	rotate(array1, 2)
	rotate(array2, 2)

	fmt.Println(array1, array2)
}

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
	return
}

func reverse(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return
}

func rotate1(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	nums = append(nums[length-k:], nums[0:length - k]...)
	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	newNums := make([]int, length)
	for i:=0; i< length;i++ {
		if i+k > length-1 {
			newNums[i+k-length] = nums[i]
		} else {
			newNums[i+k] = nums[i]
		}
	}

	copy(nums, newNums)
}

