package questionBank

import "testing"

func TestGetHint(t *testing.T) {
	secret := "1807"
	guess := "7810"
	res := getHint(secret, guess)
	if res != "1A3B" {
		t.Error("猜数字游戏算法测试1错误", res)
	}

	secret = "1123"
	guess = "0111"
	res = getHint(secret, guess)
	if res != "1A1B" {
		t.Error("猜数字游戏算法测试2错误", res)
	}

	secret = "1"
	guess = "0"
	res = getHint(secret, guess)
	if res != "0A0B" {
		t.Error("猜数字游戏算法测试3错误", res)
	}

	secret = "1"
	guess = "1"
	res = getHint(secret, guess)
	if res != "1A0B" {
		t.Error("猜数字游戏算法测试4错误", res)
	}
}
