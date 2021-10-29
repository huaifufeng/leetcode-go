/**
  题目地址：https://leetcode-cn.com/problems/self-crossing/


  解题：
  1、将投资和收益组合为一个数组，并且按照投资从小到大排序，然后获取当前资金下最大的收益值，放到总资金数，从投资组合中移除选择的项目，O(kn) 会超出时间限制
*/
package questionBank

//@todo
func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		// 第 1 类路径交叉的情况
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// 第 2 类路径交叉的情况
		if i == 4 && distance[3] == distance[1] &&
			distance[4] >= distance[2]-distance[0] {
			return true
		}

		// 第 3 类路径交叉的情况
		if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
			distance[i-1] <= distance[i-3] &&
			distance[i] >= distance[i-2]-distance[i-4] &&
			distance[i-2] > distance[i-4] {
			return true
		}
	}
	return false
}
