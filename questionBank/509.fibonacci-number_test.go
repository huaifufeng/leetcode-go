package questionBank

import "testing"

func TestFib(t *testing.T) {
	res := fib(2)
	if res != 1 {
		t.Error("斐波那契数列错误")
	}

	res = fib(3)
	if res != 2 {
		t.Error("斐波那契数列错误")
	}

	res = fib(4)
	if res != 3 {
		t.Error("斐波那契数列错误")
	}
}
