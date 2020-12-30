/**
  题目地址：https://leetcode-cn.com/problems/last-stone-weight/

  有一堆石头，每块石头的重量都是正整数。
  每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
    如果 x == y，那么两块石头都会被完全粉碎；
    如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
  最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

  <pre>
     输入：[2,7,4,1,8,1]
     输出：1
     解释：
       先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
       再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
       接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
       最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
  </pre>
	1、按照题解中说明的，逐步处理最大的两个石块，然后重新组合切片之后，继续处理最大的石块，直到剩下小于2个石块的时候，返回最后的重量
  解题：
*/
package questionBank

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	stones = smashStone(stones)

	return stones[0]
}

func smashStone(stones []int) []int {
	if len(stones) == 1 {
		return stones
	}

	if len(stones) == 0 {
		return []int{0}
	}

	for i := 0; i < 2; i++ {
		for j := 1; j < len(stones); j++ {
			if stones[i] < stones[j] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}

	subNum := stones[0] - stones[1]
	leftStones := stones[2:]
	if subNum > 0 {
		leftStones = append(leftStones, subNum)
	}

	return smashStone(leftStones)
}
