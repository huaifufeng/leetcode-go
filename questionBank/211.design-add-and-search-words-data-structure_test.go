package questionBank

import "testing"

func TestWordDictionary(t *testing.T) {
	wordD := Constructor1()
	wordD.AddWord("bad")
	wordD.AddWord("dad")
	wordD.AddWord("mad")
	if wordD.Search("pad") {
		t.Error("添加与搜索单词 - 数据结构设计 算法错误")
	}

	if !wordD.Search("bad") {
		t.Error("添加与搜索单词 - 数据结构设计 算法错误")
	}

	if !wordD.Search(".ad") {
		t.Error("添加与搜索单词 - 数据结构设计 算法错误")
	}

	if !wordD.Search("b..") {
		t.Error("添加与搜索单词 - 数据结构设计 算法错误")
	}
}
