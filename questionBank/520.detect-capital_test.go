package questionBank

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	word := "USA"
	res := detectCapitalUse(word)
	if !res {
		t.Error("检测大写字母算法测试1错误")
	}

	word = "FlaG"
	res = detectCapitalUse(word)
	if res {
		t.Error("检测大写字母算法测试2错误")
	}
}
