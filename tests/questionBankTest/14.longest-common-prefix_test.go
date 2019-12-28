package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T)  {
	intNum := questionBank.LongestCommonPrefix([]string{"flow","fl","flight"})
	if intNum != "fl" {
		t.Log(intNum)
		t.Error("公共前缀获取失败")
	}
}
