package questionBank

import "testing"

func TestFindWords(t *testing.T) {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	if len(res) != 2 || res[0] != "Alaska" {
		t.Error("键盘行算法测试1错误")
	}

	words = []string{"omk"}
	res = findWords(words)
	if len(res) != 0 {
		t.Error("键盘行算法测试2错误")
	}

	words = []string{"adsdf", "sfd"}
	res = findWords(words)
	if len(res) != 2 {
		t.Error("键盘行算法测试1错误")
	}
}
