//题目地址：https://leetcode-cn.com/problems/two-sum/
package 题库

func twoSum(num []int, target int)[]int {
	numLength := len(num)
	numMap := make(map[int]int, numLength)
	retNum := make([]int, 2)

	for i := 0; i < numLength; i++ {
		lackValue := target - num[i]
		if lackIndex, ok := numMap[num[i]]; ok {
			retNum[0] = lackIndex
			retNum[1] = i
			return retNum
		}

		numMap[lackValue] = i
	}

	return retNum
}