package questionBank

import "testing"

func TestIntegerReplacement(t *testing.T) {
	n := 8
	res := integerReplacement(n)
	if res != 3 {
		t.Error("整数替换算法测试1错误")
	}

	n = 7
	res = integerReplacement(n)
	if res != 4 {
		t.Error("整数替换算法测试2错误")
	}

	n = 4
	res = integerReplacement(n)
	if res != 2 {
		t.Error("整数替换算法测试3错误")
	}

	n = 65535
	res = integerReplacement(n)
	if res != 17 {
		t.Error("整数替换算法测试4错误")
	}
}
