/**
  题目地址：https://leetcode-cn.com/problems/keyboard-row/


  解题：
	1、建立字母的字典，循环单词列表，依次比较每个单词的位置，O(n)
*/
package questionBank

func findWords(words []string) []string {
	bMap := map[byte]uint8{
		'a': 1,
		's': 1,
		'd': 1,
		'f': 1,
		'g': 1,
		'h': 1,
		'j': 1,
		'k': 1,
		'l': 1,
		'A': 1,
		'S': 1,
		'D': 1,
		'F': 1,
		'G': 1,
		'H': 1,
		'J': 1,
		'K': 1,
		'L': 1,
		'z': 2,
		'x': 2,
		'c': 2,
		'v': 2,
		'b': 2,
		'n': 2,
		'm': 2,
		'Z': 2,
		'X': 2,
		'C': 2,
		'V': 2,
		'B': 2,
		'N': 2,
		'M': 2,
	}

	res := make([]string, 0)
	for _, word := range words {
		lastB := uint8(4)
		for i := 0; i < len(word); i++ {
			bb := bMap[word[i]]
			if lastB == 4 {
				lastB = bb
				continue
			}
			if lastB != 4 && bb != lastB {
				lastB = 5
				break
			}
		}
		if lastB != 5 {
			res = append(res, word)
		}
	}

	return res
}
