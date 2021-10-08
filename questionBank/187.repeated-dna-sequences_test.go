package questionBank

import "testing"

func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	res := findRepeatedDnaSequences(s)
	if len(res) != 2 {
		t.Error("重复的DNA序列算法测试1错误")
	}

	s = "AAAAAAAAAAAAA"
	res = findRepeatedDnaSequences(s)
	if len(res) != 1 {
		t.Error("重复的DNA序列算法测试2错误")
	}

	s = "AAAAAAAAAAA"
	res = findRepeatedDnaSequences(s)
	if len(res) != 1 {
		t.Error("重复的DNA序列算法测试3错误")
	}
}
