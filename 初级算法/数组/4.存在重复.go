//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/24/
package 数组

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