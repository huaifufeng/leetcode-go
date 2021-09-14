/**
  题目地址：https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/

  解题：
  1、三层循环这个数组，依次计算第一层点到另外两个点的距离，相等的话加2，因为可以调整顺序，O(n^3)
  2、
*/
package questionBank

func findLongestWord(s string, dictionary []string) string {
	maxStr := ""
	sSlice := []byte(s)
	for _, s2 := range dictionary {
		if len(s2) < len(maxStr) {
			continue
		}
		di := 0
		for _, v := range sSlice {
			if s2[di] == v {
				di++

				if di == len(s2) {
					if len(s2) > len(maxStr) || (len(s2) == len(maxStr) && s2 < maxStr) {
						maxStr = s2
					}
					break
				}
			}
		}
	}

	return maxStr
}
