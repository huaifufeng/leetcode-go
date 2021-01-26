/**
  题目地址：https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/

  给你一个由一些多米诺骨牌组成的列表 dominoes。
  如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
  形式上，dominoes[i] = [a, b]和dominoes[j] = [c, d]等价的前提是a==c且b==d，或是a==d 且b==c。
  在0 <= i < j < dominoes.length的前提下，找出满足dominoes[i] 和dominoes[j]等价的骨牌对 (i, j) 的数量。

  <pre>
     输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
     输出：1
  </pre>

  解题：
    1、循环进行处理，时间复杂度O(n2)  会超出时间限制
    2、将数组处理为单一值，只需要循环依次，唯一难点是增加次数
*/
package questionBank

func numEquivDominoPairs(dominoes [][]int) int {
	if len(dominoes) <= 1 {
		return 0
	}

	num := 0
	val := [100]int{}
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] < dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		v := dominoes[i][0]*10 + dominoes[i][1]
		num += val[v] //当遇到一个值和之前等价时，就需要加上之前等价的数量，这个就是新增加的等价值
		val[v]++
	}

	return num
}

func numEquivDominoPairs1(dominoes [][]int) int {
	if len(dominoes) <= 1 {
		return 0
	}

	num := 0
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				num++
			}
		}
	}

	return num
}
