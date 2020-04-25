//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/25/
package 数组

func singleNumber(nums []int) int {
	hasMap := make(map[int]int, len(nums))
	for _, val := range nums {
		if _, ok := hasMap[val]; ok {
			delete(hasMap, val)
		} else {
			hasMap[val] = 1
		}
	}

	key := 0
	for key = range hasMap {
		break
	}

	return key
}
