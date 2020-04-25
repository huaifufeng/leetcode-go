package questionBank

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	intNum := RomanToInt("MCMXCIV")
	if intNum != 1994 {
		t.Log(intNum)
		t.Error("罗马数字转换失败")
	}
}
