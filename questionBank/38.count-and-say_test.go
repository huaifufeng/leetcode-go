package questionBank

import "testing"

func TestCountAndSay(t *testing.T) {
	n := 1
	res := countAndSay(n)
	if res != "1" {
		t.Error("外观数列算法测试1错误")
	}

	n = 4
	res = countAndSay(n)
	if res != "1211" {
		t.Error("外观数列算法测试2错误", res)
	}
}
