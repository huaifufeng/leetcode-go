package questionBank

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T)  {
	intNum := LongestCommonPrefix([]string{"flow","fl","flight"})
	if intNum != "fl" {
		t.Log(intNum)
		t.Error("公共前缀获取失败")
	}
}
