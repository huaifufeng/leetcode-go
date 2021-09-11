package questionBank

import "testing"

func TestFindIntegers(t *testing.T) {
	input := 10000000
	res := findIntegers(input)
	if res != 5 {
		t.Error("不含连续1的非负整数算法测试1错误", res)
	}
}
