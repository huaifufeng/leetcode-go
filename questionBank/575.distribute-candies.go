/**
  题目地址：https://leetcode-cn.com/problems/distribute-candies/

  解题：
    1、循环整个字符串数组，依次比较当前元素和之前元素的值，进行元素值得计数。O(n)
*/
package questionBank

func distributeCandies(candyType []int) int {
	max := (len(candyType)) >> 1
	candyMap := make(map[int]struct{}, max)
	for _, u := range candyType {
		candyMap[u] = struct{}{}
		if len(candyMap) == max {
			return max
		}
	}

	return len(candyMap)
}
