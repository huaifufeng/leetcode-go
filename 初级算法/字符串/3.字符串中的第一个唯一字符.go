//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/34/
package 字符串


func firstUniqChar(s string) int {
	length := len(s)
	hasMap := make(map[uint8]int)
	deleteMap := make(map[uint8]int)

	for i := 0; i < length; i++ {
		if _,ok := deleteMap[s[i]]; ok {
			continue
		}
		hasMap[s[i]] = i
		for j := i+1; j < length; j++ {
			if _, ok := hasMap[s[j]]; ok {
				delete(hasMap, s[i])
				deleteMap[s[i]] = 1
				break
			}
		}

		if _, ok := hasMap[s[i]]; ok {
			return i
		}
	}

	return -1
}

func firstUniqChar1(s string) int {
	length := len(s)
	hasMap := make(map[uint8]int, 26)

	for i:=0;i<length;i++ {
		hasMap[s[i]]++
	}

	for i:=0; i<length; i++ {
		if hasMap[s[i]] == 1 {
			return i
		}
	}

	return -1
}