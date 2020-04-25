package questionBank

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	value := LengthOfLongestSubstring("abba")
	if value != 2 {
		t.Log(value)
		t.Error("方法验证失败")
	}
}
