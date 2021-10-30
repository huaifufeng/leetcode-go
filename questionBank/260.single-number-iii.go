/**
  题目地址：https://leetcode-cn.com/problems/single-number-iii/


  解题：
	1、循环列表，存储数据的map值，根据map=1获取需要的元素 O(n)
*/
package questionBank

func singleNumber(nums []int) []int {
	numMaps := make(map[int]uint8)
	for _, num := range nums {
		numMaps[num]++
		if numMaps[num] == 2 {
			delete(numMaps, num)
		}
	}

	res := make([]int, 0)
	for i, _ := range numMaps {
		res = append(res, i)
	}

	return res
}
