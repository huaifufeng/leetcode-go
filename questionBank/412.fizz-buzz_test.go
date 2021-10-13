package questionBank

import "testing"

func TestFizzBuzz(t *testing.T) {
	n := 3
	res := fizzBuzz(n)
	if res[2] != "Fizz" {
		t.Error("Fizz Buzz算法测试1错误")
	}

	n = 5
	res = fizzBuzz(n)
	if res[4] != "Buzz" {
		t.Error("Fizz Buzz算法测试2错误")
	}

	n = 15
	res = fizzBuzz(n)
	if res[14] != "FizzBuzz" {
		t.Error("Fizz Buzz算法测试3错误")
	}
}
