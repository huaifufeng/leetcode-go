/**
  题目地址：https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk/

  解题：
  1、最多循环2遍。选循环一遍获取每个索引位置之前数据的和如果找到直接返回，否则对总和取余之后，再循环一遍就可以 O(n)
*/
package questionBank

func chalkReplacer(chalk []int, k int) int {
	usedMap := make(map[int]int)
	for i, i2 := range chalk {
		if usedMap[i-1]+i2 > k {
			return i
		}
		usedMap[i] = usedMap[i-1] + i2
	}

	left := k % usedMap[len(chalk)-1]
	for i, _ := range chalk {
		if usedMap[i] > left {
			return i
		}
	}

	return -1
}
