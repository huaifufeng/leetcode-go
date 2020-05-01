package questionBank

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	value := lengthOfLongestSubstring("abba")
	if value != 2 {
		t.Error("获取最长子串方法验证失败")
	}
}
