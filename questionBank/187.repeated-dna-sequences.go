/**
  题目地址：https://leetcode-cn.com/problems/repeated-dna-sequences/

  解题：
  1、循环字符串，获取每个长度为10的子串的数量，放到字典中，在获取大于1的子串，O(nL) n字符串长度 L 子串长度
*/
package questionBank

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	sMap := make(map[string]uint8)
	bSlice := []byte(s)
	for i := 10; i <= len(bSlice); i++ {
		subStr := string(bSlice[i-10 : i])
		sMap[subStr]++
	}

	res := make([]string, 0)
	for s2, u := range sMap {
		if u > 1 {
			res = append(res, s2)
		}
	}

	return res
}
