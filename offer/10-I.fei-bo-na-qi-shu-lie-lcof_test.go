package offer

import "testing"

func TestFib(t *testing.T) {
	n := 2
	res := fib(n)
	if res != 1 {
		t.Error("斐波那契数列算法测试1错误", res)
	}

	n = 5
	res = fib(n)
	if res != 5 {
		t.Error("斐波那契数列算法测试2错误", res)
	}

	n = 43
	res = fib(n)
	if res != 433494437 {
		t.Error("斐波那契数列算法测试3错误", res)
	}
}
