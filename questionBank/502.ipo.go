/**
  题目地址：https://leetcode-cn.com/problems/ipo/

  给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
  你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

  <pre>
     给定 nums = [2, 7, 11, 15], target = 9

     因为 nums[0] + nums[1] = 2 + 7 = 9
     所以返回 [0, 1]
  </pre>

  解题：
  1、将投资和收益组合为一个数组，并且按照投资从小到大排序，然后获取当前资金下最大的收益值，放到总资金数，从投资组合中移除选择的项目，O(kn) 会超出时间限制
*/
package questionBank

import "sort"

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	cp := make([][]int, 0)

	//获取指定投资的最大收益
	for i, i2 := range capital {
		cp = append(cp, []int{i2, profits[i]})
	}

	sort.Slice(cp, func(i, j int) bool {
		return cp[i][0] < cp[j][0]
	})

	for i := 0; i < k; i++ {
		max := 0
		index := 0
		for i2, cpInfo := range cp {
			if w < cpInfo[0] {
				break
			}

			if cpInfo[1] > max {
				max = cpInfo[1]
				index = i2
			}
		}
		cp = append(cp[:index], cp[index+1:]...)
		w += max
		if len(cp) == 0 {
			break
		}
	}

	return w
}
