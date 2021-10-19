/**
  题目地址：https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/

  解题：
  1、
*/
package questionBank

type WordDictionary struct {
	Words map[string]uint8
}

func Constructor1() WordDictionary {
	words := WordDictionary{
		Words: make(map[string]uint8),
	}

	return words
}

func (this *WordDictionary) AddWord(word string) {
	this.Words[word] = 1
}

func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.Words[word]; ok {
		return true
	}

	sMap := make(map[int]byte)
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			sMap[i] = word[i]
		}
	}

	for key, _ := range this.Words {
		if len(key) != len(word) {
			continue
		}

		if len(sMap) == 0 {
			return true
		}

		flag := true
		for i, v := range sMap {
			if key[i] != v {
				flag = false
				break
			}
		}

		if flag {
			return true
		}
	}

	return false
}
