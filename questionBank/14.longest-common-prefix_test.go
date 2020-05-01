package questionBank

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	intNum := longestCommonPrefix([]string{"flow", "fl", "flight"})
	if intNum != "fl" {
		t.Error("获取公共前缀方法失败")
	}
}
