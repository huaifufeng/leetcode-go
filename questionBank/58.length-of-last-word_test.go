package questionBank

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	res := lengthOfLastWord("Hello World")
	if res != 5 {
		t.Error("最后单词长度算法错误")
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	res := lengthOfLastWord2("Hello World")
	if res != 5 {
		t.Error("最后单词长度算法2错误")
	}
}
