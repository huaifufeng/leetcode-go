//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/35/
package 字符串

//异位词：字母相同，只是位置不同
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, 26)
	tMap := make(map[rune]int, 26)

	for _, val := range s {
		sMap[val]++
	}

	for _, val := range t {
		tMap[val]++
	}

	for key := range sMap {
		if tMap[key] != sMap[key] {
			return false
		}
	}

	return true
}