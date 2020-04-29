package questionBank

import "testing"

func TestIsHappyTrue(t *testing.T) {
	n := 1111111
	res := isHappy(n)

	if res != true {
		t.Error("判断快乐数方法错误")
	}
}

func TestIsHappyFalse(t *testing.T) {
	n := 11
	res := isHappy(n)
	if res != false {
		t.Error("判断快乐数方法错误")
	}
}

func TestIsHappy2True(t *testing.T) {
	n := 1111111
	res := isHappy2(n)

	if res != true {
		t.Error("判断快乐数方法错误")
	}
}

func TestIsHappy2False(t *testing.T) {
	n := 11
	res := isHappy2(n)
	if res != false {
		t.Error("判断快乐数方法错误")
	}
}
