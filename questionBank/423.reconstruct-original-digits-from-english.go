/**
  题目地址：https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/

  解题：
  1、最简单的方法就是循环整个数组，然后依次将数组中的值相加，然后得到想要的结果。但是这种在数量很大的时候，会非常耗费时间，时间复杂度O(n²)
  2、我们循环数组，将数组的值和索引以 值：索引的方式放入一个map中，这样在后面的循环中，只需要判断目标值-当期值=差值 这个差值在不在map中，
  只要在map中存在，就说明找到了目前的组合(因为这里确定有且只有一个组合存在，那么找到一个组合就完成任务了)，时间复杂度O(n)
*/
package questionBank

import (
	"strconv"
	"strings"
)

func originalDigits(s string) string {
	rMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		rMap[s[i]]++
	}

	nnMaps := make(map[int]int)
	nnMaps[0] = rMap['z']
	nnMaps[2] = rMap['w']
	nnMaps[6] = rMap['x']
	nnMaps[4] = rMap['u']
	nnMaps[8] = rMap['g']

	nnMaps[7] = rMap['s'] - nnMaps[6]
	nnMaps[3] = rMap['h'] - nnMaps[8]
	nnMaps[5] = rMap['f'] - nnMaps[4]

	nnMaps[1] = rMap['o'] - nnMaps[0] - nnMaps[2] - nnMaps[4]
	nnMaps[9] = rMap['i'] - nnMaps[5] - nnMaps[6] - nnMaps[8] - nnMaps[9]

	res := ""
	for i := 0; i < 10; i++ {
		if nnMaps[i] > 0 {
			res += strings.Repeat(strconv.Itoa(i), nnMaps[i])
		}
	}

	return res
}
