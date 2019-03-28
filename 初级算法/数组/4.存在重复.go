package main

import "fmt"

func main() {
	nums := []int{1,2,3,1}

	result := containsDuplicate(nums)

	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	hasMap := make(map[int]int, len(nums))
	for _,val := range nums {
		if _,ok := hasMap[val];ok {
			return true
		} else {
			hasMap[val] = 1
		}
	}

	return false
}