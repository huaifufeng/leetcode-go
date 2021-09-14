package questionBank

import "testing"

func TestFindLongestWord(t *testing.T) {
	str := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	res := findLongestWord(str, dictionary)
	if res != "apple" {
		t.Error("通过删除字母匹配到字典里最长单词算法测试1错误", res)
	}

	str = "abpcplea"
	dictionary = []string{"a", "b", "c"}
	res = findLongestWord(str, dictionary)
	if res != "a" {
		t.Error("通过删除字母匹配到字典里最长单词算法测试2错误", res)
	}

	str = "abce"
	dictionary = []string{"abe", "abc"}
	res = findLongestWord(str, dictionary)
	if res != "abc" {
		t.Error("通过删除字母匹配到字典里最长单词算法测试3错误", res)
	}
}
