//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/29/
package 数组

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