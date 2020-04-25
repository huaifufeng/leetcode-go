package questionBank

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	reverseInt := PalindromeNumber(1)
	if reverseInt != true {
		t.Log(reverseInt)
		t.Error("正数反转失败")
	}
}
