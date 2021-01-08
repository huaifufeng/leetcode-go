/**
  题目地址：https://leetcode-cn.com/problems/number-of-provinces/

  有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
  省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
  给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
  返回矩阵中 省份 的数量。

  <pre>
     输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
	 输出：2
  </pre>

  <pre>
	 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
	 输出：3
  </pre>

  解题：
    1、
*/
package questionBank

//@todo 这里没有计算后面造成的前面几项合并问题的
func findCircleNum(isConnected [][]int) int {
	num := 0
	checked := make(map[int]uint8, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		flag := false
		for j := i; j < len(isConnected[i]); j++ {
			val := isConnected[i][j]
			if val == 0 {
				continue
			}
			_, ok := checked[j]
			if ok {
				continue
			}

			checked[j] = 1
			flag = true
		}

		if flag {
			num++
		}
	}

	return num
}
