/**
  题目地址：https://leetcode-cn.com/problems/missing-number/

  解题：
  1、存储已经存在的值的字典，获取不存在的值
*/
package questionBank

func missingNumber(nums []int) int {
	numMap := make(map[int]uint8)
	for _, num := range nums {
		numMap[num] = 1
	}

	for i := 0; i <= len(nums); i++ {
		if numMap[i] == 0 {
			return i
		}
	}

	return -1
}
